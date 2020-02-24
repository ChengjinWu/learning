package xlsx

import (
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"regexp"
	"testing"
)

var (
	numRe, _ = regexp.Compile(`\d{12,}`)
)

func TestRe(t *testing.T) {
	fmt.Println(numRe.MatchString("6010350612961557"))
}
func TestXlsx(t *testing.T) {
	excelFileName := "/Users/chengjinwu/Documents/ad_server/src/learning/xlsx/所有联盟展现2.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	var pv, uv int
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) == 7 {
				if row.Cells[0].String() == "20190415" {
					var flag bool
					if row.Cells[3].String() == "tencent" {
						flag = true
					} else {
						if numRe.Match([]byte(row.Cells[4].String())) {
							flag = true
							fmt.Printf("%s\n", printRow(row))
						} else {
							//t.Errorf("fail:%s", printRow(row))
						}

					}
					if flag {
						v1, err := row.Cells[5].Int()
						if err != nil {
							t.Errorf("%s", printRow(row))
							t.Error(err)
						}
						v2, err := row.Cells[6].Int()
						if err != nil {
							t.Errorf("%s", printRow(row))
							t.Error(err)
						}
						pv += v1
						uv += v2
					}
				}

			} else {
				fmt.Printf("%s\n", printRow(row))
			}
		}
	}
	fmt.Printf("result:%d,%d", pv, uv)
}

func printRow(row *xlsx.Row) []byte {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%5d,", len(row.Cells)))
	for _, cell := range row.Cells {
		text := cell.String()
		buf.WriteString(fmt.Sprintf("%20s,", text))
	}
	return buf.Bytes()
}
