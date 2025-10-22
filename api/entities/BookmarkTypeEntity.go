package entities

type BookmarkType struct {
	BookmarkTypeId   int    `gorm:"column:bookmark_type_id;primaryKey;not null" json:"bookmark_type_id"`
	BookmarkTypeName string `gorm:"bookmark_type_name" json:"bookmark_type_name"`
}

func (*BookmarkType) TableName() string {
	return "mtr_bookmark_type"
}
