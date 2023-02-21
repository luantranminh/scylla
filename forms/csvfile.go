package forms

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"scylla/models"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type CsvFile struct {
	Keywords []string `form:"csvfile" valid:"Required"`
}

func (r *CsvFile) Valid(v *validation.Validation) {

	fmt.Println("we have ", r.Keywords, len(r.Keywords))

	if len(r.Keywords) > 100 {
		v.SetError("Keyword", "Your quota is 100 keywords")
		return
	}

	for _, keyword := range r.Keywords {
		if len(keyword) > 200 {
			v.SetError("Keyword", keyword+" is too long (200 is the limit)")
			break
		}
	}

}

func Parse(file multipart.File) CsvFile {
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		logs.Error(err)
	}

	csvStruct := CsvFile{
		Keywords: make([]string, len(csvLines)),
	}

	for i, line := range csvLines {
		csvStruct.Keywords[i] = strings.Join(line, " ")
	}

	return csvStruct
}

func (r *CsvFile) Save() error {
	validation := validation.Validation{}

	valid, err := validation.Valid(r)

	if err != nil {
		logs.Error(err)
		return err
	}

	if !valid {
		for _, err := range validation.Errors {
			return err
		}
	}

	mKeywords := make([]models.Keyword, len(r.Keywords))
	for i, keyword := range r.Keywords {
		mKeywords[i] = models.Keyword{
			Content: keyword,
			Status:  models.New,
		}
	}

	_, err = models.AddKeywords(mKeywords)

	return err
}
