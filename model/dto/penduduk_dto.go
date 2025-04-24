package dto

type RegistrasiInput struct {
	Nama          string
	JenisKelamin  string
	TempatLahir   string
	TanggalLahir  string
	KodeProvinsi  string
	KodeKota      string
	KodeKecamatan string
}

type UpdatePendudukInput struct {
	Nama        string
	TempatLahir string
}
