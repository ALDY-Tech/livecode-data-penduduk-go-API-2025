package repository

import (
	"Building-the-API-Layer/model"
	"database/sql"
)

type PendudukRepository interface {
	CreatePenduduk(penduduk *model.Penduduk) error
	GetPendudukByNik(nik string) (*model.Penduduk, error)
	GetPendudukById(id int) (*model.Penduduk, error)
	UpdatePenduduk(penduduk *model.Penduduk) error
	DeletePenduduk(id int) error
	GetAllPenduduk() ([]model.Penduduk, error)
}

type pendudukRepository struct {
	db *sql.DB
}

func(p *pendudukRepository) CreatePenduduk(penduduk *model.Penduduk) error {
	// insert into penduduk (nik, nama, jenis_kelamin, tempat_lahir, tanggal_lahir, kode_provinsi, kode_kota, kode_kecamatan) values (?, ?, ?, ?, ?, ?, ?, ?)
	_, err := p.db.Exec("insert into penduduk (nik, nama, jenis_kelamin, tempat_lahir, tanggal_lahir, kode_provinsi, kode_kota, kode_kecamatan) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		penduduk.Nik,
		penduduk.Nama,
		penduduk.JenisKelamin,
		penduduk.TempatLahir,
		penduduk.TanggalLahir,
		penduduk.KodeProvinsi,
		penduduk.KodeKota,
		penduduk.KodeKecamatan)

	if err != nil {
		return err
	}
	return nil
}

func(p *pendudukRepository) GetPendudukByNik(nik string) (*model.Penduduk, error) {
	// select * from penduduk where nik = ?
	var penduduk model.Penduduk
	err := p.db.QueryRow("select * from penduduk where nik = ?", nik).Scan(
		&penduduk.ID,
		&penduduk.Nik,
		&penduduk.Nama,
		&penduduk.JenisKelamin,
		&penduduk.TempatLahir,
		&penduduk.TanggalLahir,
		&penduduk.KodeProvinsi,
		&penduduk.KodeKota,
		&penduduk.KodeKecamatan,
		&penduduk.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &penduduk, nil
	}

func(p *pendudukRepository) GetPendudukById(id int) (*model.Penduduk, error) {
	// select * from penduduk where id = ?
	var penduduk model.Penduduk
	err := p.db.QueryRow("select * from penduduk where id = ?", id).Scan(
		&penduduk.ID,
		&penduduk.Nik,
		&penduduk.Nama,	
		&penduduk.JenisKelamin,
		&penduduk.TempatLahir,	
		&penduduk.TanggalLahir,
		&penduduk.KodeProvinsi,
		&penduduk.KodeKota,
		&penduduk.KodeKecamatan,
		&penduduk.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &penduduk, nil
	}

func(p *pendudukRepository) UpdatePenduduk(penduduk *model.Penduduk) error {
	// update penduduk set nama = ?, tempat_lahir = ? where id = ?
	_, err := p.db.Exec("update penduduk set nama = ?, tempat_lahir = ? where id = ?",
		penduduk.Nama,
		penduduk.TempatLahir,
		penduduk.ID)

	if err != nil {
		return err
	}
	return nil
}

func(p *pendudukRepository) DeletePenduduk(id int) error {
	// delete from penduduk where id = ?
	_, err := p.db.Exec("delete from penduduk where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func(p *pendudukRepository) GetAllPenduduk() ([]model.Penduduk, error) {
	// select * from penduduk
	rows, err := p.db.Query("select * from penduduk")
	if err != nil {
		return nil, err
	}

	var Penduduk []model.Penduduk
	for rows.Next() {
		var penduduk model.Penduduk
		err := rows.Scan(
			&penduduk.ID,
			&penduduk.Nik,
			&penduduk.Nama,
			&penduduk.JenisKelamin,
			&penduduk.TempatLahir,
			&penduduk.TanggalLahir,
			&penduduk.KodeProvinsi,
			&penduduk.KodeKota,
			&penduduk.KodeKecamatan,
			&penduduk.CreatedAt)
		if err != nil {
			return nil, err
		}
		Penduduk = append(Penduduk, penduduk)
	}
	return Penduduk, nil
}

func NewPendudukRepository(db *sql.DB) PendudukRepository {
	repo := new(pendudukRepository)
	repo.db = db
	return repo
}