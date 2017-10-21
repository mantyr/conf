package tmpl

import (
	//    "github.com/mantyr/conf"
	"html/template"
	//    "path/filepath"
	//    "sync"
	"fmt"
	"strings"
)

type PaginationLink struct {
	url       string
	page      int64
	link_type string
	is_active bool
}

func NewPaginationLink() (p *PaginationLink) {
	p = new(PaginationLink)
	return
}

func (p *PaginationLink) GetPage() int64 {
	return p.page
}

func (p *PaginationLink) SetPage(page int64) *PaginationLink {
	p.page = page
	return p
}

func (p *PaginationLink) SetUrl(url string) *PaginationLink {
	p.url = url
	return p
}

func (p *PaginationLink) GetAddress() string {
	return strings.Replace(p.url, "{page}", fmt.Sprintf("%v", p.page-1), -1)
}

func (p *PaginationLink) IsType(link_type string) bool {
	if p.link_type == link_type {
		return true
	}
	return false
}

func (p *PaginationLink) IsActive() bool {
	return p.is_active
}

func (p *PaginationLink) SetActive() {
	p.is_active = true
}

func (p *PaginationLink) SetType(link_type string) *PaginationLink {
	p.link_type = link_type
	return p
}

func Pagination(url string, page int64, page_max int64, address string) (r template.HTML, err error) {
	data := struct {
		Url      string
		Page     int64
		Page_max int64
		Links    []*PaginationLink
	}{
		url,
		page,
		page_max,
		nil,
	}
	var i int64
	page++
	for i = 1; i <= page_max; i++ {
		l := NewPaginationLink()
		if i == page {
			l.SetActive()
		}
		l.SetUrl(url)
		l.SetType("link")
		l.SetPage(i)
		data.Links = append(data.Links, l)

		if page-i > 5 {
			i = page - 4
			l := NewPaginationLink()
			l.SetType("...")

			data.Links = append(data.Links, l)
		} else if (i-page) == 3 && (page_max-i) > 2 {
			i = page_max - 1
			l := NewPaginationLink()
			l.SetType("...")

			data.Links = append(data.Links, l)
		}
	}

	return RenderHTML(data, address)
}
