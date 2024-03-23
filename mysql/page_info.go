package mysql

import (
	"context"
	"go_study/structure"
	"gorm.io/gorm"
)

type (
	PageInfo struct {
		Page     int    `excel:"page"`
		Title    string `excel:"title"`
		Href     string `excel:"href"`
		Src      string `excel:"src"`
		InfoType string
		gorm.Model
	}
	PageInfoList []*PageInfo
)

func (m *PageInfo) TableName() string {
	return "page_info"
}

type (
	PageInfoDao struct {
		db *gorm.DB
	}
	IPageInfoDao interface {
		CreatePageInfo(ctx context.Context, baseInfo *PageInfo) error
		CreatePageInfoList(ctx context.Context, list PageInfoList) error
		Search(ctx context.Context, q structure.QueryPageInfoParam) (PageInfoList, int64, error)
	}
)

func NewPageInfoDao(db *gorm.DB) IPageInfoDao {
	db.AutoMigrate(&PageInfo{})
	return &PageInfoDao{
		db: db,
	}
}

func (d *PageInfoDao) CreatePageInfo(ctx context.Context, baseInfo *PageInfo) error {
	return d.db.Create(baseInfo).Error
}

func (d *PageInfoDao) CreatePageInfoList(ctx context.Context, list PageInfoList) error {
	return d.db.Create(list).Error
}

func (d *PageInfoDao) Search(ctx context.Context, q structure.QueryPageInfoParam) (PageInfoList, int64, error) {
	var (
		err   error
		count int64
		list  = make(PageInfoList, 0)
	)
	db := d.db
	if q.InfoType != "" {
		db = db.Where("info_type = ?", q.InfoType)
	}
	if q.Title != "" {
		db = db.Where("title like ?", "%"+q.Title+"%")
	}
	if q.TitleEqual != "" {
		db = db.Where("title = ?", q.Title)
	}
	if q.Src != "" {
		db = db.Where("src like ?", "%"+q.Src+"%")
	}
	if q.Href != "" {
		db = db.Where("href like ?", "%"+q.Href+"%")
	}
	if q.Size > 0 {
		db = db.Limit(q.Size).Offset(q.GetOffSet())
	}
	err = db.Find(&list).Error
	db.Count(&count)
	return list, count, err
}
