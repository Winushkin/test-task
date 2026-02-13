//Package report creates reports of tsv files' records in pdf files
package report

import(
	"codeberg.org/go-pdf/fpdf"
	"file-manager/internal/parser"
	"reflect"
	"fmt"
)

const fontFilePath = "./internal/report/fonts/TikTokSans_18pt_Condensed-Bold.ttf"
const eightSpaces = "        "
func CreateReportsFromFile(records []parser.Record, dirPath string) error{
	for _, record := range records{
		err := createReport(record, dirPath)
		if err != nil{
			return fmt.Errorf("createReport: %w", err)
		}
	}
	return nil
}

func createReport(record parser.Record, dirPath string) (error){
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Sans", "", fontFilePath)
	pdf.SetFont("Sans", "", 10)

	t := reflect.TypeOf(record)
	v := reflect.ValueOf(record)
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		name := field.Name
		info := field.Tag.Get("info")

		pdf.Cell(100, 20, fmt.Sprintf("%s (%s)", name, info))
		pdf.Ln(5)

		var displayValue string
		if !v.Field(i).IsZero(){
			displayValue = fmt.Sprintf("%v", v.Field(i).Interface())
		}else{
			displayValue = "-"
		}
		pdf.Cell(40, 20, fmt.Sprintf("%s%s", eightSpaces, displayValue))
		pdf.Ln(10)
	}

	err := pdf.OutputFileAndClose(fmt.Sprintf("%s/%s.pdf", dirPath, record.UnitGUID))
	if err != nil {
		return fmt.Errorf("OutputFileAndClose: %w", err)
	}
	return nil
}