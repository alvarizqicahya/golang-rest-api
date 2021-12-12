package model

import (
	"golang-rest-api/app"
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	// gorm.Model
	Id        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Nama      string         `gorm:"type:varchar(255)" json:"nama"`
	NoTelp    string         `gorm:"type:varchar(255)" json:"no_telp"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type MahasiswaModel struct{}

// type MahasiswaCreate struct {
// 	Nama   string `json:"nama"`
// 	NoTelp string `json:"no_telp"`
// }

func (M *MahasiswaModel) ShowAll() ([]Mahasiswa, error) {
	var mhs []Mahasiswa
	err := app.InitDB.Find(&mhs).Error
	return mhs, err
}

func (M *MahasiswaModel) Find(id string) (Mahasiswa, error) {
	var mhs Mahasiswa
	err := app.InitDB.Where("id = ?", id).First(&mhs).Error
	return mhs, err
}

func (M *MahasiswaModel) Save(data *Mahasiswa) error {
	err := app.InitDB.Save(data).Error
	return err
}

func (M *MahasiswaModel) Delete(id string) error {
	return app.InitDB.Where("id = ?", id).Delete(&Mahasiswa{}).Error
}
