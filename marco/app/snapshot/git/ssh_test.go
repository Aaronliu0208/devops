package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"
)

var remoteRepo string = "git@yunludev3.htyunlu.com:shanyou/marco-snapshot.git"

func TestSSHKeyload(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	wd := filepath.Dir(filename)
	s := fmt.Sprintf("%s/test/id_rsa", wd)
	localRepo := fmt.Sprintf("%s/test/localrepo", wd)
	sshKey, err := ioutil.ReadFile(s)
	if err != nil {
		t.Fatal(err)
	}

	pubkey, err := ssh2.NewPublicKeys("git", []byte(sshKey), "1234.asd")
	if err != nil {
		t.Fatal(err)
	}

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteRepo},
	})
	_, err = rem.List(&git.ListOptions{
		Auth: pubkey,
	})
	if err != nil && err != transport.ErrEmptyRemoteRepository {
		t.Fatal(err)
	}
	if err == transport.ErrEmptyRemoteRepository {
		fmt.Println("empty repo")
	}
	wt := osfs.New(localRepo)
	dot, _ := wt.Chroot(".git")
	fs := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	repo, err := git.Init(fs, wt)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// os.RemoveAll(localRepo + "/.git")
	}()
	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteRepo},
	})

	if err != nil {
		t.Fatal(err)
	}

	list, err := repo.Remotes()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range list {
		fmt.Println(r)
	}

	w, err := repo.Worktree()
	err = filepath.Walk(localRepo,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, ".git") {
				return nil
			}
			if !info.IsDir() {
				fmt.Println(path, info.Size())
				w.Add(path)
			}

			return nil
		})
	status, err := w.Status()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status)

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "shanyou",
			Email: "shanyou@htyunwang.com",
			When:  time.Now(),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	// Prints the current HEAD to verify that all worked well.

	obj, err := repo.CommitObject(commit)
	fmt.Println(obj)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Push(&git.PushOptions{
		Auth: pubkey,
	})

	if err != nil {
		t.Fatal(err)
	}
}
