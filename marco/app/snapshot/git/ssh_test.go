package git

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

var remoteURL string = "git@yunludev3.htyunlu.com:shanyou/marco-snapshot.git"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func getAuth() (transport.AuthMethod, error) {
	_, filename, _, _ := runtime.Caller(0)
	wd := filepath.Dir(filename)
	s := fmt.Sprintf("%s/test/id_rsa", wd)
	sshKey, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}

	pubkey, err := ssh2.NewPublicKeys("git", []byte(sshKey), "1234.asd")
	if err != nil {
		return pubkey, err
	}

	return pubkey, err
}

func getRepoDir() string {
	_, filename, _, _ := runtime.Caller(0)
	wd := filepath.Dir(filename)
	return fmt.Sprintf("%s/test/localrepo", wd)
}

func CreateRemote() *git.Remote {
	return git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteURL},
	})
}
func init() {
	rand.Seed(time.Now().UnixNano())
	localRepo := getRepoDir()
	os.RemoveAll(localRepo)
	os.MkdirAll(localRepo, 0777)
}

func destroy() {
	localRepo := getRepoDir()
	os.RemoveAll(localRepo)
	os.MkdirAll(localRepo, 0777)
}

func checkIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestSSHKeyload(t *testing.T) {
	defer destroy()
	localRepo := getRepoDir()
	pubkey, err := getAuth()
	checkIfError(t, err)

	rem := CreateRemote()
	_, err = rem.List(&git.ListOptions{
		Auth: pubkey,
	})
	if err != nil && err != transport.ErrEmptyRemoteRepository {
		t.Fatal(err)
	}

	var isEmpty bool = false
	if err == transport.ErrEmptyRemoteRepository {
		fmt.Println("empty repo")
		isEmpty = true
	}
	wt := osfs.New(localRepo)
	dot, _ := wt.Chroot(".git")
	fs := filesystem.NewStorage(dot, cache.NewObjectLRUDefault())
	repo, err := git.Init(fs, wt)
	checkIfError(t, err)

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{remoteURL},
	})
	checkIfError(t, err)

	list, err := repo.Remotes()
	checkIfError(t, err)

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
	checkIfError(t, err)
	fmt.Println(status)

	commit, err := w.Commit("example go-git commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "shanyou",
			Email: "shanyou@htyunwang.com",
			When:  time.Now(),
		},
	})

	checkIfError(t, err)
	// Prints the current HEAD to verify that all worked well.

	obj, err := repo.CommitObject(commit)
	fmt.Println(obj)
	checkIfError(t, err)

	if !isEmpty {
		// fetch and merge
		err = repo.Fetch(&git.FetchOptions{
			RemoteName: "origin",
			Auth:       pubkey,
		})

		checkIfError(t, err)

		err = w.Pull(&git.PullOptions{RemoteName: "origin",
			Auth: pubkey})
		checkIfError(t, err)

	}

	err = repo.Push(&git.PushOptions{
		Auth: pubkey,
	})

	checkIfError(t, err)
}

func TestCheckout(t *testing.T) {
	defer destroy()
	rem := CreateRemote()
	auth, err := getAuth()
	checkIfError(t, err)
	refs, err := rem.List(&git.ListOptions{
		Auth: auth,
	})

	fmt.Printf("%v\n", refs)
	if err != nil && err != transport.ErrEmptyRemoteRepository {
		t.Fatal(err)
	}
	localRepo := getRepoDir()
	r, err := git.PlainClone(localRepo, false, &git.CloneOptions{
		Auth:       auth,
		NoCheckout: true,
		Progress:   os.Stdout,
		URL:        remoteURL,
	})
	checkIfError(t, err)
	fmt.Printf("%v\n", r)
	wt, err := r.Worktree()
	checkIfError(t, err)
	rnd := randString(10)
	ioutil.WriteFile(filepath.Join(localRepo, "test.log"), []byte(rnd), 0644)
	status, err := wt.Status()
	checkIfError(t, err)
	for n, _ := range status {
		wt.Add(n)
	}
	commit, err := wt.Commit("commit with:"+rnd, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "shanyou",
			Email: "shanyou@htyunwang.com",
			When:  time.Now(),
		},
	})

	checkIfError(t, err)
	obj, err := r.CommitObject(commit)
	fmt.Println(obj)
	checkIfError(t, err)
	err = r.Push(&git.PushOptions{
		Auth: auth,
	})

	checkIfError(t, err)

}
