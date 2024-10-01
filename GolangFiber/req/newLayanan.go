package req

type NewLayanan struct {
	Id          uint   `gorm:"primaryKey" json:"idLayanan"`
	NamaLayanan string `gorm:"type:varchar(100)" json:"namaLayanan"`
	Deskripsi   string `gorm:"type:longtext" json:"deskripsi"`
}
