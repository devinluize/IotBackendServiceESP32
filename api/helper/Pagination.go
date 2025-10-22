package helper

import (
	"fmt"
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	SortOf     string      `json:"sort_of"`
	SortBy     string      `json:"sort_by"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return p.GetPage() * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetSortOf() string {
	if p.SortOf == "" {
		p.SortOf = "asc"
	}
	return p.SortOf
}

func (p *Pagination) GetSortBy() string {
	return p.SortBy
}
func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	var sort string = ""
	if pagination.GetSortBy() != "" {
		sort = pagination.GetSortBy() + " " + pagination.GetSortOf()
	}
	//db.Model()
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := 0
	if totalRows != 0 {
		totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	}
	pagination.TotalPages = totalPages
	ofset := pagination.GetOffset()
	limit := pagination.GetLimit()

	fmt.Println(ofset, limit)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(sort)
	}
}
