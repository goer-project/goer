package paginator

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Total       int64       `json:"total"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Data        interface{} `json:"data"`
}

type Paginator struct {
	PerPage  int
	Page     int
	Offset   int
	Total    int64
	LastPage int
	Sort     string
	Order    string

	query *gorm.DB
	ctx   *gin.Context
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, perPage int) Pagination {
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage)

	err := p.query.Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		return Pagination{}
	}

	return Pagination{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		LastPage:    p.LastPage,
		Total:       p.Total,
		Data:        data,
	}
}

func (p *Paginator) initProperties(perPage int) {

	p.PerPage = p.getPerPage(perPage)

	p.Order = p.ctx.DefaultQuery("order", "desc")
	p.Sort = p.ctx.DefaultQuery("sort", "id")

	p.Total = p.getTotalCount()
	p.LastPage = p.getLastPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p Paginator) getPerPage(perPage int) int {
	queryPerPage := p.ctx.Query("per_page")
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	if perPage <= 0 {
		perPage = 15
	}

	return perPage
}

func (p Paginator) getCurrentPage() int {
	page := cast.ToInt(p.ctx.Query("page"))
	if page <= 0 {
		page = 1
	}

	if p.LastPage == 0 {
		return 0
	}

	if page > p.LastPage {
		return p.LastPage
	}

	return page
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p Paginator) getLastPage() int {
	if p.Total == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.Total) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}

	return int(nums)
}
