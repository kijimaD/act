package report

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
	"act/config"
	"act/scrape"
)

type Report struct {
	in     []scrape.Repo
	config config.Config
}

func NewReport(in scrape.Scrape, config config.Config) *Report {
	return &Report{
		in:     in.Repos,
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
			panic(err)
		}
		defer file.Close()
		output = file
	default:
		panic("output type error. stdout | file")
	}

	r.header(output)
	r.repoTable(output)

	return r
}

func (r *Report) header(output io.Writer) {
	str := fmt.Sprintf("# central\n")
	str += fmt.Sprintf("%v public repos\n\n", len(r.in))
	output.Write([]byte(str))
}

func (r *Report) repoTable(output io.Writer) {
	table := tablewriter.NewWriter(output)
	table.SetHeader(r.repoHeaders())
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for _, repo := range r.in {
		table.Append(r.repoContent(repo))
	}

	table.Render()
}

func (r *Report) repoHeaders() []string {
	return []string{"Name", "Description", "Language", "Forks", "Star", "Commit"}
}

func (r *Report) repoContent(repo scrape.Repo) []string {
	return []string{repo.MdLink(), repo.Description, repo.Language, strconv.Itoa(repo.ForksCount), strconv.Itoa(repo.StargazersCount), strconv.Itoa(repo.CommitCount)}
}
