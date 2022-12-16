package model

type DRankRights struct {
	Pk       int32 `gorm:"pk" json:"pk"`
	RightsPk int32 `gorm:"rights_pk" json:"rights_pk"` // 权益PK
	RankPk   int32 `gorm:"rank_pk" json:"rank_pk"`     // 头衔PK
}

func (*DRankRights) TableName() string {
	return "d_rank_rights"
}
