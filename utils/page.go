package utils

//分页工具
type Page struct {
	PageNow      int64
	PageSize     int64
	PrevPage     bool
	NextPage     bool
	TotalRecNum  int64
	TotalPageNum int64
	PageContent  interface{}
	StartIndex   int64
	EndIndex     int64
}

func NewPage() *Page {
	page := &Page{
		PageNow:  1,
		PageSize: 5,
	}
	return page
}

func (this *Page) GetStartIndex() int64 {
	//在前一页的基础上再加一条
	return this.PageSize*(this.PageNow-1) + 1 // size:10 pageno:3   21
}

func (this *Page) GetEndIndex() int64 {
	recNums := this.PageSize * this.PageNow
	// total:25  7/page , 取第4页      22-25
	if recNums > this.TotalRecNum {
		return this.TotalRecNum
	}
	return recNums
}
