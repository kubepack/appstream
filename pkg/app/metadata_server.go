package app

import (
	"fmt"
	"strings"

	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"github.com/appscode/appstream/pkg/apiserver/endpoints"
	"github.com/appscode/appstream/pkg/app/docker"
	"github.com/appscode/appstream/pkg/app/git"
	"github.com/appscode/appstream/pkg/util"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(api.RegisterMetadataServer, &MetadataServer{})
	endpoints.ProxyServerEndpoints.Register(api.RegisterMetadataHandlerFromEndpoint)
}

type MetadataServer struct{}

var _ api.MetadataServer = &MetadataServer{}

func (*MetadataServer) Get(ctx context.Context, req *api.MetadataGetRequest) (*api.MetadataGetResponse, error) {
	resp := &api.MetadataGetResponse{}

	if strings.EqualFold(req.Type, "docker") {
		if md, err := docker.GetMetadata(req.Name, req.Registry); err != nil {
			resp.Status = util.NewStatusFromError(err)
			return resp, nil
		} else {
			resp.Metdata = &api.MetadataGetResponse_Docker{
				Docker: md,
			}
		}
	} else if strings.EqualFold(req.Type, "git") {
		if md, err := git.GetMetadata(req.Name, req.Registry); err != nil {
			resp.Status = util.NewStatusFromError(err)
			return resp, nil
		} else {
			resp.Metdata = &api.MetadataGetResponse_Git{
				Git: md,
			}
		}
	} else {
		resp.Status = util.NewStatusFromError(fmt.Errorf("Regsitry type %v not supported.", req.Type))
		return resp, nil
	}
	resp.Status = util.NewStatusOK()
	return resp, nil
}
