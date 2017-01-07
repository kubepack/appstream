// Code generated by protoc-gen-grpc-gateway-cors
// source: metdata.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportMetadataCorsPatterns returns an array of grpc gatway mux patterns for
// Metadata service to enable CORS.
func ExportMetadataCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Metadata_Get_0,
	}
}
