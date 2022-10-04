package report

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"strconv"
	"act/config"
	"act/scrape"
	"sort"
	"strings"
)

type Report struct {
	in     []scrape.Repo
	config config.Config
	langMap map[string]langRecord
	commitCount int
	repoCount int
}

func NewReport(in scrape.Scrape, config config.Config) *Report {
	sort.Slice(in.Repos, func(i, j int) bool { return in.Repos[i].Language < in.Repos[j].Language })

	langMap := make(map[string]langRecord)
	for _, repo := range in.Repos {
		newRepo := langMap[repo.Language].repoC + 1
		newCommit := langMap[repo.Language].commitC + repo.CommitCount

		langMap[repo.Language] = langRecord{
			repoC: newRepo,
			commitC: newCommit,
		}
	}

	commitCount := 0
	for _, repo := range in.Repos {
		commitCount += repo.CommitCount
	}

	return &Report{
		in:     in.Repos,
		config: config,
		langMap: langMap,
		commitCount: commitCount,
		repoCount: len(in.Repos),
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
	r.langTable(output)
	r.repoTable(output)

	return r
}

func (r *Report) blankLine(output io.Writer) {
	output.Write([]byte("\n"))
}

func (r *Report) header(output io.Writer) {
	str := fmt.Sprintf("# central\n\n")
	str += fmt.Sprintf("- %v public repos\n", r.repoCount)
	str += fmt.Sprintf("- %v commits\n", r.commitCount)
	str += fmt.Sprintf("\n")
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
	r.blankLine(output)
}

type langRecord struct {
	repoC int
	commitC int
}

func (r *Report) langTable(output io.Writer) {
	table := tablewriter.NewWriter(output)
	table.SetHeader(r.langHeaders())
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for k, v := range r.langMap {
		table.Append(r.langContent(k, v))
	}

	table.Render()
	r.blankLine(output)
}

func (r *Report) langHeaders() []string {
	return []string{
		"Lang",
		"Repo",
		"%",
		"Commit",
		"%",
	}
}

func (r *Report) langContent(name string, l langRecord) []string {
	return []string{
		fmt.Sprintf("%v %v", r.dotImage(name), name),
		fmt.Sprintf("%v", l.repoC),
		fmt.Sprintf("%v%%", int(float64(l.repoC) / float64(r.repoCount) * 100)),
		fmt.Sprintf("%v", l.commitC),
		fmt.Sprintf("%v%%", int(float64(l.commitC) / float64(r.commitCount) * 100)),
	}
}

func (r *Report) dotImage(name string) string {
	newName := strings.Replace(name, " ", "-", -1)
	return fmt.Sprintf("![%v](https://raw.githubusercontent.com/kijimaD/maru/main/images/dot/%v.svg)", name, newName)
}

func (r *Report) repoHeaders() []string {
	return []string{
		"Name",
		"Desc",
		"Lang",
		"Commit",
		"Star",
		"Fork",
	}
}

func (r *Report) repoContent(repo scrape.Repo) []string {
	return []string{
		repo.MdLink(),
		repo.Description,
		fmt.Sprintf("%v %v", r.dotImage(repo.Language), repo.Language),
		strconv.Itoa(repo.CommitCount),
		strconv.Itoa(repo.StargazersCount),
		strconv.Itoa(repo.ForksCount),
	}
}
