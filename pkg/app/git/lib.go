package git

import (
	"os"
	"strings"

	api "github.com/appscode/appstream/pkg/apis/app/v1beta1"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

const origin = "origin"

func GetMetadata(name, reg string) (*api.GitMetadata, error) {
	r, err := git.NewFilesystemRepository(cacheDir(name))
	if err != nil {
		return nil, err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: origin,
		URL:  name,
	})
	if err != nil {
		return nil, err
	}

	err = r.Pull(&git.PullOptions{
		RemoteName: origin,
	})

	refs, err := r.References()
	if err != nil {
		return nil, err
	}

	md := &api.GitMetadata{
		Name:     name,
		Branches: make([]string, 0),
		Notes:    make([]string, 0),
		Tags:     make([]string, 0),
	}
	err = refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic
		// references, the HEAD
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}
		if ref.IsBranch() {
			md.Branches = append(md.Branches, ref.Name().Short())
		} else if ref.IsNote() {
			md.Notes = append(md.Notes, ref.Name().Short())
		} else if ref.IsTag() {
			md.Tags = append(md.Tags, ref.Name().Short())
		}
		return nil
	})

	err = r.DeleteRemote(origin)
	if err != nil {
		return nil, err
	}

	return md, nil
}

func cacheDir(name string) string {
	name = strings.TrimSuffix(name, ".git")
	name = strings.Replace(name, "://", "-", -1)
	name = strings.Replace(name, "@", "-", -1)
	name = strings.Replace(name, ":", "-", -1)
	name = strings.Replace(name, "/", "-", -1)

	path := os.TempDir() + "/" + name
	os.MkdirAll(path, 0755)
	return path
}
