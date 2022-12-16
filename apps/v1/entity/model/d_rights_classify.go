package model

type DRightsClassify struct {
	Pk          int32  `gorm:"pk" json:"pk"`
	Name        string `gorm:"name" json:"name"`                 // 名称
	SupperiorPk int32  `gorm:"supperior_pk" json:"supperior_pk"` // 上级PK
	Description string `gorm:"description" json:"description"`   // 描述
}

func (*DRightsClassify) TableName() string {
	return "d_rights_classify"
}
