package gh

import (
	ghttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"time"
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/v47/github"
	"os"
	"net/http"
	"act/config"
	"github.com/go-git/go-git/v5"
	"path/filepath"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const botName = "github-actions[bot]"
const botEmail = "github-actions[bot]@users.noreply.github.com"

type gh struct {
	Client *github.Client
	Config config.Config
}

func New(c config.Config) *gh {
	client := github.NewClient(Login())

	return &gh{
		Client: client,
		Config: c,
	}
}

func Login() *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_TOKEN")},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	return tc
}

func (gh *gh) User() *github.User {
	// Passing the empty string will fetch the authenticated user.
	user, _, err := gh.Client.Users.Get(context.Background(), "")
	if err != nil {
		panic(err)
	}
	return user
}

// コミット
// 作業に影響しないようにインメモリGitで操作するのがベストだが、ファイル操作の方法がわからなかったのでカレントディレクトリで作業する
func (gh *gh) Commit() *gh {
	if gh.Config.IsCommit {
		directory := "./"

		// Opens an already existing repository.
		r, err := git.PlainOpen(directory)
		if err != nil {
			panic(err)
		}
		w, err := r.Worktree()
		if err != nil {
			panic(err)
		}

		// Adds the new file to the staging area.
		_, fname := filepath.Split(gh.Config.OutPath)
		_, err = w.Add(fname)
		fmt.Println(gh.Config.OutPath)
		if err != nil {
			panic(err)
		}

		// We can verify the current status of the worktree using the method Status.
		status, _ := w.Status()
		fmt.Println(status)

		// commit
		commit, err := w.Commit("commit by act [ci skip]", &git.CommitOptions{
			Author: &object.Signature{
				Name:  botName,
				Email: botEmail,
				When:  time.Now(),
			},
		})
		obj, err := r.CommitObject(commit)

		fmt.Println(obj)
	}
	return gh
}

// support only HTTPS push
func (gh *gh) Push() {
	if gh.Config.IsPush {
		path := "./"

		r, err := git.PlainOpen(path)
		if err != nil {
			panic(err)
		}

		err = r.PushContext(
			context.Background(),
			&git.PushOptions{
				Progress: os.Stdout,
				Auth: &ghttp.BasicAuth{
					Username: gh.Config.UserId,
					Password: os.Getenv("GH_TOKEN"),
				},
			})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			panic(err)
		}
		fmt.Println("Pushed")
	}
}
