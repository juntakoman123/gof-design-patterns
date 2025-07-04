package main

import (
	"fmt"
	"os"
)

type Report interface {
	FetchData() string
	FormatData(data string) string
	Output(formatted string)
}

type CSVReport struct{}

func (r *CSVReport) FetchData() string {
	return "name,age\nAlice,30\nBob,25"
}

func (r *CSVReport) FormatData(data string) string {
	return "[CSV形式で出力]\n" + data
}

func (r *CSVReport) Output(formatted string) {
	fmt.Println(formatted)
}

type FileReport struct{}

func (r *FileReport) FetchData() string {
	return "title: Monthly Report\ndata: 100 entries"
}

func (r *FileReport) FormatData(data string) string {
	return "[ファイル出力形式]\n" + data
}

func (r *FileReport) Output(formatted string) {
	file, err := os.Create("report.txt")
	if err != nil {
		fmt.Println("ファイル作成失敗:", err)
		return
	}
	defer file.Close()

	file.WriteString(formatted)
	fmt.Println("ファイルに書き出しました")
}

type ReportGenerator struct {
	impl Report
}

func (rg *ReportGenerator) GenerateReport() {
	data := rg.impl.FetchData()
	formatted := rg.impl.FormatData(data)
	rg.impl.Output(formatted)
}

func main() {
	// CSVレポートを生成
	csvGen := &ReportGenerator{impl: &CSVReport{}}
	csvGen.GenerateReport()

	// ファイルレポートを生成
	fileGen := &ReportGenerator{impl: &FileReport{}}
	fileGen.GenerateReport()
}
