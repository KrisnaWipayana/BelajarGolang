package model

type User struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	Role       int         `gorm:"type:tinyint(1)" json:"role"`
	Nama       string      `gorm:"type:varchar(200)" json:"nama"`
	Email      string      `gorm:"type:varchar(100)" json:"email"`
	Password   string      `gorm:"type:varchar(100)" json:"password"`
	Pemesanans []Pemesanan `gorm:"foreignKey:UserID" json:"pemesanans"`
}

type Kamar struct {
	Id        uint       `gorm:"primaryKey" json:"idKamar"`
	NamaKamar string     `gorm:"type:varchar(100)" json:"namaKamar"`
	Harga     float64    `gorm:"type:decimal(10,2)" json:"harga"` // Tambahkan harga
	Fotos     []Foto     `gorm:"foreignKey:KamarID" json:"fotos"`
	NoKamars  []No_Kamar `gorm:"foreignKey:KamarID" json:"noKamars"`
	Layanans  []Layanan  `gorm:"many2many:kamar_layanans;" json:"layanans"`
}

type Foto struct {
	Id       uint   `gorm:"primaryKey" json:"idFoto"`
	NamaFoto string `gorm:"type:varchar(100)" json:"namaFoto"`
	KamarID  uint   `gorm:"index" json:"kamarId"`
}

type Layanan struct {
	Id          uint    `gorm:"primaryKey" json:"idLayanan"`
	NamaLayanan string  `gorm:"type:varchar(100)" json:"namaLayanan"`
	Kamars      []Kamar `gorm:"many2many:kamar_layanans;" json:"kamars"`
}

type No_Kamar struct {
	Id      uint `gorm:"primaryKey" json:"idNoKamar"`
	KamarID uint `gorm:"index" json:"kamarId"`
}

type Pemesanan struct {
	Id        uint      `gorm:"primaryKey" json:"idPemesanan"`
	UserID    uint      `gorm:"index" json:"idUser"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	KamarID   uint      `gorm:"index" json:"idKamar"`
	Kamar     Kamar     `gorm:"foreignKey:KamarID" json:"kamar"`
	NamaKamar string    `json:"namaKamar"`                                     // Nama kamar dari Kamar struct
	NamaUser  string    `json:"namaUser"`                                      // Nama user dari User struct
	Layanans  []Layanan `gorm:"many2many:pemesanan_layanans;" json:"layanans"` // Layanan yang dipilih
}
