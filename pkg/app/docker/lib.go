package docker

import (
	//"github.com/docker/distribution/digest"
	//"github.com/docker/distribution/manifest"
	//"github.com/docker/libtrust"
	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"github.com/heroku/docker-registry-client/registry"
)

func GetMetdata(name, reg string) (*api.DockerMetadata, error) {
	url := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous

	hub, err := registry.New(url, username, password)
	if err != nil {
		return nil, err
	}

	tags, err := hub.Tags(name)
	if err != nil {
		return nil, err
	}
	/*
		for _, tag := range tags {
			manifest, err := hub.Manifest("appscode/kubed", tag)
			if err != nil {
				log.Fatal(err)
			}
			bb, err := manifest.MarshalJSON()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(bb))
		}
	*/
	return &api.DockerMetadata{
		Name: name,
		Tags: tags,
	}, nil
}
