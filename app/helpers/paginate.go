package helpers

import (
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

const PAGE_SIZE = 10

func Paginate(request *http.Request) func(database *gorm.DB) *gorm.DB {
	return func(database *gorm.DB) *gorm.DB {
		page := GetPage(request)

		offset := int((page - 1) * PAGE_SIZE)

		return database.Offset(offset).Limit(PAGE_SIZE)
	}
}

func GetPage(request *http.Request) int64 {
	page, err := strconv.Atoi(request.URL.Query().Get("page"))

	if page == 0 || err != nil {
		page = 1
	}

	return int64(page)
}

func NewPaginator(output interface{}, database *gorm.DB, request *http.Request, pageSize int64) Paginator {
	var totalItems int64

	database.Count(&totalItems)

	totalPages := int64(math.Ceil(float64(totalItems) / float64(pageSize)))

	Paginate(request)(database).Find(output)

	currentPage := GetPage(request)

	if currentPage > totalPages {
		currentPage = totalPages
	}

	var previousPage *int64
	var previous = currentPage - 1

	if previous > 0 {
		previousPage = &previous
	} else {
		previousPage = nil
	}

	var nextPage *int64
	var next = currentPage + 1

	if currentPage < totalPages && next <= totalPages {
		nextPage = &next
	} else {
		nextPage = nil
	}

	return Paginator{
		TotalItems:   totalItems,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		LastPage:     totalPages,
		CurrentPage:  currentPage,
		Items:        output,
	}
}

type Paginator struct {
	TotalItems   int64
	NextPage     *int64
	PreviousPage *int64
	LastPage     int64
	CurrentPage  int64
	Items        interface{}
}
