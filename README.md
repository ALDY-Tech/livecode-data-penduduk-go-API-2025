# Building API Layer

## Request

Sample Request Register

```json
{
  "nama": "John Doe",
  "jenis_kelamin": "Laki-laki",
  "tempat_lahir": "Jakarta",
  "tanggal_lahir": "2000-10-01",
  "kode_provinsi": "18",
  "kode_kota": "1809",
  "kode_kecamatan": "180905"
}
```

Sample Request Update

```json
{
  "nama": "John Doe",
  "tempat_lahir": "Jakarta"
}
```

## Response

Sample Single Response

```json
{
  "message": "Register successful",
  "data": {
    "id": 1,
    "nik": "1809050110000001",
    "nama": "John Doe",
    "jenis_kelamin": "Laki-laki",
    "tempat_lahir": "Jakarta",
    "tanggal_lahir": "2000-10-01",
    "kode_provinsi": "18",
    "kode_kota": "1809",
    "kode_kecamatan": "180905",
    "created_at": "2000-10-01 10:00"
  }
}
```

Sample Single Response

```json
{
  "message": "Get all data successful",
  "data": [
    {
      "id": 1,
      "nik": "1809050110000001",
      "nama": "John Doe",
      "jenis_kelamin": "Laki-laki",
      "tempat_lahir": "Jakarta",
      "tanggal_lahir": "2000-10-01",
      "kode_provinsi": "18",
      "kode_kota": "1809",
      "kode_kecamatan": "180905",
      "created_at": "2000-10-01 10:00"
    },
    {
      "id": 2,
      "nik": "1809054110000001",
      "nama": "Maria Jona",
      "jenis_kelamin": "Laki-laki",
      "tempat_lahir": "Jakarta",
      "tanggal_lahir": "2000-10-01",
      "kode_provinsi": "18",
      "kode_kota": "1809",
      "kode_kecamatan": "180905",
      "created_at": "2000-10-01 10:00"
    }
  ]
}
```
