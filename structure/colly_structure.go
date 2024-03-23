package structure

type QueryPageInfoParam struct {
	ListQuery
	Page       int
	Title      string
	TitleEqual string
	Href       string
	Src        string
	InfoType   string
}
