package model

type MetaAssetSuffix struct {
	Pk          int32  `gorm:"pk" json:"pk"`
	Name        string `gorm:"name" json:"name"`               // 名称
	Description string `gorm:"description" json:"description"` // 描述
	Scene       string `gorm:"scene" json:"scene"`             // 场景
}

func (*MetaAssetSuffix) TableName() string {
	return "meta_asset_suffix"
}
