package pdf

import (
	"log"

	"github.com/signintech/gopdf"
)

func Sample() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("loma", "assets/font/loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(156, 197, 140)
	pdf.Cell(nil, "สวัสดี")
	pdf.WritePdf("./output/hello.pdf")
}
