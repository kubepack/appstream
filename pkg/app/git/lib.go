package git

import (
	"fmt"
	"os"

	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func GetMetadata(name, reg string) (*api.GitMetadata, error) {
	r, err := git.NewFilesystemRepository(os.TempDir())
	if err != nil {
		return nil, err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URL:  "https://github.com/tidwall/buntdb.git",
	})
	if err != nil {
		return nil, err
	}

	list, err := r.Remotes()
	if err != nil {
		return nil, err
	}

	for _, r := range list {
		fmt.Println(r)
	}

	err = r.Pull(&git.PullOptions{
		RemoteName: "origin",
	})

	refs, err := r.References()
	if err != nil {
		return nil, err
	}

	md := &api.GitMetadata{
		Name:     name,
		Branches: make([]string, 0),
		Tags:     make([]string, 0),
	}
	err = refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic
		// references, the HEAD
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}
		if ref.IsBranch() {
			md.Branches = append(md.Branches, ref.Name().String())
		} else if ref.IsTag() {
			md.Tags = append(md.Tags, ref.Name().String())
		}
		return nil
	})

	err = r.DeleteRemote("origin")
	if err != nil {
		return nil, err
	}

	return md, nil
}
