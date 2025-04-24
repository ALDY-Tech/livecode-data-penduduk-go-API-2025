package utils

import (
	"fmt"
	"strings"
	"time"
)

func GenerateNIKLogic(kodeProv, kodeKota, kodeKec string, tglLahir time.Time, jenisKelamin string) (string, error) {

	nikWilayah := kodeProv + kodeKota[2:4] + kodeKec[4:6]
	day := tglLahir.Day()
	month := int(tglLahir.Month())
	year := tglLahir.Year() % 100

	jenisKelaminUpper := strings.ToUpper(jenisKelamin)
	if jenisKelaminUpper == "PEREMPUAN" {
		day += 40
	} else if jenisKelaminUpper != "LAKI-LAKI" {
		return "", fmt.Errorf("jenis kelamin tidak valid: '%s'", jenisKelamin)
	}

	// Format tanggal DD-MM-YY
	nikTanggal := fmt.Sprintf("%02d%02d%02d", day, month, year)

	// Nomor Urut (4 digit) - Hardcoded / random
	nikUrut := "0001"

	nik := nikWilayah + nikTanggal + nikUrut

	if len(nik) != 16 {
		return "", fmt.Errorf("NIK tidak boleh lebih dari 16 digit, panjang: %d", len(nik))
	}

	return nik, nil
}
