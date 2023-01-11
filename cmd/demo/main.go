package main

import (
	"fmt"

	"github.com/captain-corgi/go-doc-generator/pkg/word"
)

type (
	AIATemplate struct {
		CompanyName string
		CompanyCode string
		IssueDate   string
		TotalAmount string
		AccountNo   string
		Date        string
		Phone       string
	}
)

func main() {
	docxTemplate := word.NewDocxLegionZverTemplate()
	template, err := docxTemplate.OpenTemplate("assets/template/aia_template.docx")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 20; i++ {
		data := AIATemplate{
			CompanyName: fmt.Sprintf("My Company No %d", i),
			CompanyCode: "123456",
			IssueDate:   "10-Jan-2023",
			TotalAmount: "1000 Baht",
			AccountNo:   fmt.Sprintf("H-0092-%d", i),
			Date:        "10-Jan-2023",
			Phone:       "Tel. 034 295 5286",
		}
		if err := template.RenderTemplate(data); err != nil {
			fmt.Println(err)
			return
		}
		if err := template.Save(fmt.Sprintf("output/aia_%s.docx", data.AccountNo)); err != nil {
			fmt.Println(err)
			return
		}
		if err := template.ResetTemplate(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Success")
	}
}
