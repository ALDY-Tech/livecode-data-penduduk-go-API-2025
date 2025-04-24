package model

import "time"

type Penduduk struct {
	ID            int
	Nik           string
	Nama          string
	JenisKelamin  string
	TempatLahir   string
	TanggalLahir  time.Time
	KodeProvinsi  string
	KodeKota      string
	KodeKecamatan string
	CreatedAt     time.Time
}
