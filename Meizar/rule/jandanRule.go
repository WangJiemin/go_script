package rule

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type JandenRule struct{}

func (p *JandenRule) UrlRule() (url string) {
	return "http://jandan.net/ooxx/"
}

func (p *JandenRule) PageRule(currentPage int) (page string) {
	return "page-" + strconv.Itoa(currentPage)
}

func (p *JandenRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
		if img, exist := s.Attr("href"); exist {
			f(img)
		}
	})
}
