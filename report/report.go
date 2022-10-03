package report

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
	"act/config"
	"act/scrape"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	ghttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"time"
	"path/filepath"
)

type Report struct {
	in     []scrape.Repo
	out    string
	config config.Config
}

func NewReport(in scrape.Scrape, config config.Config) *Report {
	return &Report{
		in:     in.Repos,
		out:    "",
		config: config,
	}
}

func (r *Report) Execute() *Report {
	// 出力分岐を別の関数にしたいが、出力されない(空白のファイルになる)
	var output io.Writer

	switch r.config.OutType {
	case "stdout":
		output = os.Stdout
	case "file":
		file, err := os.Create(r.config.OutPath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		output = file
	default:
		panic("output type error. stdout | file")
	}

	table := tablewriter.NewWriter(output)
	// table := tablewriter.NewWriter(r.selectOutput())
	table.SetHeader(r.headers())
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for _, repo := range r.in {
		table.Append(r.content(repo))
	}

	table.Render()
	return r
}

func (r *Report) headers() []string {
	return []string{"Name", "Description", "Language", "Forks", "Star", "Commit"}
}

func (r *Report) content(repo scrape.Repo) []string {
	return []string{repo.MdLink(), repo.Description, repo.Language, strconv.Itoa(repo.ForksCount), strconv.Itoa(repo.StargazersCount), strconv.Itoa(repo.CommitCount)}
}

// コミット
// 作業に影響しないようにインメモリGitで操作するのがベストだが、ファイル操作の方法がわからなかったのでカレントディレクトリで作業する
func (report *Report) Commit() *Report {
	if report.config.IsCommit {
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
		_, fname := filepath.Split(report.config.OutPath)
		_, err = w.Add(fname)
		fmt.Println(report.config.OutPath)
		if err != nil {
			panic(err)
		}

		// We can verify the current status of the worktree using the method Status.
		status, _ := w.Status()
		fmt.Println(status)

		// commit
		commit, err := w.Commit("commit by act", &git.CommitOptions{
			Author: &object.Signature{
				Name:  report.config.User.Name,
				Email: report.config.User.Email,
				When:  time.Now(),
			},
		})
		obj, err := r.CommitObject(commit)

		fmt.Println(obj)
	}
	return report
}

func (report *Report) Push() {
	if report.config.IsPush {
		path := "./"

		r, err := git.PlainOpen(path)
		if err != nil {
			panic(err)
		}

		err = r.PushContext(
			context.Background(),
			&git.PushOptions{
				RemoteName: "https",
				Auth: &ghttp.BasicAuth{
					Username: "kijimaD",
					Password: os.Getenv("GH_TOKEN"),
				},
			})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			panic(err)
		}
		fmt.Println("Pushed")
	}
}
