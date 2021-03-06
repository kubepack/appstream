// Code generated by protoc-gen-grpc-gateway-cors
// source: version.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportVersionCorsPatterns returns an array of grpc gatway mux patterns for
// Version service to enable CORS.
func ExportVersionCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Version_Get_0,
	}
}
