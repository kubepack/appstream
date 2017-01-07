package app

import (
	"github.com/appscode/api/dtypes"
	"github.com/appscode/api/version"
	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"github.com/appscode/appstream/pkg/apiserver/endpoints"
	"github.com/appscode/appstream/pkg/util"
	v "github.com/appscode/go/version"
	"golang.org/x/net/context"
)

func init() {
	endpoints.GRPCServerEndpoints.Register(api.RegisterVersionServer, &VersionServer{})
	endpoints.ProxyServerEndpoints.Register(api.RegisterVersionHandlerFromEndpoint)
}

type VersionServer struct{}

var _ api.VersionServer = &VersionServer{}

func (*VersionServer) Get(ctx context.Context, req *dtypes.VoidRequest) (*api.VersionGetResponse, error) {
	resp := &api.VersionGetResponse{}
	resp.Version = &version.Version{
		Version:         v.Version.Version,
		VersionStrategy: v.Version.VersionStrategy,
		Os:              v.Version.Os,
		Arch:            v.Version.Arch,
		CommitHash:      v.Version.CommitHash,
		GitBranch:       v.Version.GitBranch,
		GitTag:          v.Version.GitTag,
		CommitTimestamp: v.Version.CommitTimestamp,
		BuildTimestamp:  v.Version.BuildTimestamp,
		BuildHost:       v.Version.BuildHost,
		BuildHostOs:     v.Version.BuildHostOs,
		BuildHostArch:   v.Version.BuildHostArch,
	}
	resp.Status = util.NewStatusOK()
	return resp, nil
}
