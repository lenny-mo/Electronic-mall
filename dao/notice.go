package dao

import (
	"context"
	"eletronicMall/model"

	"gorm.io/gorm"
)

// noticeDAO 用于数据库的交互
type NoticeDAO struct {
	*gorm.DB
}

func NewNoticeDAO(ctx context.Context) *NoticeDAO {
	return &NoticeDAO{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDAO {
	return &NoticeDAO{DB: db}
}

// GetNoticeById 根据提供的 ID 获取通知（Notice）记录。
// 参数:
//
//	id (uint): 要查询的通知的唯一标识符。
//
// 返回值:
//
//	notice (*model.Notice): 如果找到匹配的通知，将返回包含通知信息的结构体指针。
//	err (error): 如果发生任何错误，将返回一个非空的错误对象，否则为 nil。
func (dao *NoticeDAO) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return
}
