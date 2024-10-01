package req

type CreateKamar struct {
	NamaKamar      string  `gorm:"type:varchar(100)" json:"namaKamar"`
	Harga          float64 `gorm:"type:decimal(10,2)" json:"harga"`
	DeskripsiKamar string  `gorm:"type:longtext" json:"deskripsiKamar"`
}
