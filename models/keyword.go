package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type Status string

var (
	New       Status = "new"
	Inprocess Status = "in_process"
	Done      Status = "done"
)

type Keyword struct {
	Base

	Content string `orm:"size(200)"`
	Status  Status `orm:"size(20);default(new)"`
}

func (u *Keyword) TableName() string {
	return "keywords"
}

func init() {
	orm.RegisterModel(new(Keyword))
}

// AddKeyword insert a new Keyword into database and returns
// last inserted Id on success.
func AddKeyword(m *Keyword) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddKeywords(keywords []Keyword) ([]Keyword, error) {
	if len(keywords) < 1 {
		return nil, nil
	}

	o := orm.NewOrm()

	qs := o.QueryTable(keywords[0].TableName())
	pr, _ := qs.PrepareInsert()
	for i, keyword := range keywords {
		id, err := pr.Insert(&keyword)

		fmt.Println(keyword)
		if err != nil {
			return nil, err
		}

		keywords[i].Id = id
	}

	pr.Close() // Don't forget to close the statement

	return keywords, nil
}
