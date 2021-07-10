package lib

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const HhshGuessUrl = "https://lab.magiconch.com/api/nbnhhsh/guess"

type Query struct {
	Text     string
	textList []string
}

type WordList []struct {
	Name  string   `json:"name"`
	Trans []string `json:"trans"`
}

func NewQuery(text string) *Query {
	return &Query{
		Text:     text,
		textList: makePayloadText(text),
	}
}

func makePayloadText(text string) []string {
	r := regexp.MustCompile("([a-zA-z0-9]{2,})+")
	return r.FindAllString(text, -1)
}

func (q Query) DoQuery() WordList {
	contentType := "application/json"
	data := fmt.Sprintf(`{"text":"%s"}`, strings.Join(q.textList, ","))
	resp, err := http.Post(HhshGuessUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}

	defer resp.Body.Close()

	var t WordList
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		fmt.Println("err=", err)
		return nil
	}
	return t

}

func PrintTable(words WordList) {
	var data [][]string
	for _, wordItem := range words {
		for _, item := range wordItem.Trans {
			dataItem := []string{wordItem.Name, item}
			data = append(data, dataItem)
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Fucking Words", "HHSH"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}
