package app

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	gort "runtime"
	"strings"
	"time"

	"github.com/appscode/appstream/cmd/appstream/app/options"
	"github.com/appscode/appstream/pkg/apiserver/endpoints"
	"github.com/appscode/go/runtime"
	"github.com/appscode/log"
	gwrt "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type apiServer struct {
	Port       int
	CACertFile string
	CertFile   string
	KeyFile    string

	GRPCServer *grpc.Server
	ProxyMux   *gwrt.ServeMux
}

func (s *apiServer) ListenAndServe() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)

	// We first match the connection against HTTP2 fields. If matched, the
	// connection will be sent through the "grpcl" listener.
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	// Otherwise, we match it againts HTTP1 methods. If matched,
	// it is sent through the "httpl" listener.
	httpl := m.Match(cmux.Any())

	// Then we used the muxed listeners.
	go s.ServeGRPC(grpcl)
	if s.CACertFile == "" && s.CertFile == "" && s.KeyFile == "" {
		go s.ServeHTTP(httpl)
	} else {
		go s.ServeHTTPS(httpl)
	}

	log.Fatalln(m.Serve())
}

func (s *apiServer) ServeGRPC(l net.Listener) {
	defer runtime.HandleCrash()
	for _, endpoint := range endpoints.GRPCServerEndpoints {
		log.Infoln("Registering server:", reflect.TypeOf(endpoint.Server))
		endpoints.RegisterGRPC(endpoint.RegisterFunc, s.GRPCServer, endpoint.Server)
	}

	log.Infoln("[GRPCSERVER] Starting gRPC Server at port", s.Port)
	log.Fatalln("[GRPCSERVER] gRPC Server failed:", s.GRPCServer.Serve(l))
}

var grpcDialOptions = []grpc.DialOption{grpc.WithInsecure()}

func (s *apiServer) ServeHTTP(l net.Listener) {
	defer runtime.HandleCrash()
	for _, endpoint := range endpoints.ProxyServerEndpoints {
		log.Infoln("Registering endpoint:", funcName(endpoint.RegisterFunc))
		endpoints.RegisterProxy(endpoint.RegisterFunc, context.Background(), s.ProxyMux, fmt.Sprintf("127.0.0.1:%d", s.Port), grpcDialOptions)
	}

	log.Infoln("[PROXYSERVER] Sarting Proxy Server at port", s.Port)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      s.ProxyMux,
	}
	log.Fatalln("[PROXYSERVER] Proxy Server failed:", srv.Serve(l))
}

func (s *apiServer) ServeHTTPS(l net.Listener) {
	// Load certificates.
	certificate, err := tls.LoadX509KeyPair(s.CertFile, s.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		Ref:
		 - https://blog.cloudflare.com/exposing-go-on-the-internet/
		 - http://www.bite-code.com/2015/06/25/tls-mutual-auth-in-golang/
		 - http://www.hydrogen18.com/blog/your-own-pki-tls-golang.html
	*/
	tlsConfig := &tls.Config{
		Certificates:             []tls.Certificate{certificate},
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
		SessionTicketsDisabled:   true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			// tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		ClientAuth: tls.VerifyClientCertIfGiven,
	}
	if s.CACertFile != "" {
		caCert, err := ioutil.ReadFile(s.CACertFile)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.ClientCAs = caCertPool
	}

	// Create TLS listener.
	tlsl := tls.NewListener(l, tlsConfig)

	// Serve HTTP over TLS.
	s.ServeHTTP(tlsl)
}

func Run(cfg *options.Config) {
	cfgBytes, _ := json.Marshal(cfg)
	log.Infoln("Configuration:", string(cfgBytes))

	server := &apiServer{
		Port:       cfg.APIPort,
		CACertFile: cfg.CACertFile,
		CertFile:   cfg.CertFile,
		KeyFile:    cfg.KeyFile,
		GRPCServer: grpc.NewServer(),
		ProxyMux:   gwrt.NewServeMux(),
	}
	go server.ListenAndServe()

	go func() {
		log.Infoln(http.ListenAndServe(fmt.Sprintf(":%d", cfg.PprofPort), nil))
	}()
}

func funcName(i interface{}) string {
	name := gort.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
