package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Siswa struct {
	gorm.Model
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nama         string `json:"nama"`
	Umur         string `json:"umur"`
	JenisKelamin string `json:"jenis_kelamin"`
	Email        string `json:"email"`
	NoTelp       string `json:"no_telp"`
}

func (Siswa) TableName() string {
	return "Siswa"
}

func (s *Siswa) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s.Password), 14)
	if err != nil {
		log.Fatal("Gagal Hashing")
		return err
	}
	s.Password = string(bytes)
	return nil
}
