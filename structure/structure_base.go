package structure

type ListQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (l *ListQuery) GetOffSet() int {
	if l.Page < 1 {
		l.Page = 1
	}
	return l.Size * l.Page
}
