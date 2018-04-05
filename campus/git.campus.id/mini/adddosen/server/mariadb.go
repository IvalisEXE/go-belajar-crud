package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDosen = `insert into Dosen(Kd_Dosen, Nama_Dosen, Status, CreateBy, CreateOn)values (?,?,?,?,?)`
	selectDosenByKdDosen  = `select Kd_Dosen, Nama_Dosen, Status from Dosen where KdDosen = ?`
	selectDosen       = `select Kd_Dosen, Nama_Dosen, status from Dosen where Status = '1'`
	updateDosen       = `update Dosen set Kd_Dosen=?, Nama_Dosen=?, Status=? where Kd_Dosen=?`
	selectDosenByKeterangan = `select Kd_Dosen, Nama_Mahasiswa, Keterangan, Status from Dosen where Nama_Dosen=?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddDosen(dosen Dosen) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDosen, dosen.KdDosen, dosen.NamaDosen, dosen.Keterangan, OnAdd, dosen.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}


func (rw *dbReadWriter) ReadDosenByKdDosen(kddosen int32) (Dosen, error) {
	dosen := Dosen{KdDosen: dosen}
	err := rw.db.QueryRow(selectDosenByKdDosen, kddosen).Scan(&dosen.NamaDosen, &dosen.Status)

	if err != nil {
		return Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func (rw *dbReadWriter) ReadMahasiswa() (Mahasiswas, error) {
	mahasiswa := Mahasiswas{}
	rows, _ := rw.db.Query(selectMahasiswa)
	defer rows.Close()
	for rows.Next() {
		var m Mahasiswa
		err := rows.Scan(&m.Nim, &m.NamaMahasiswa, &m.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return mahasiswa, err
		}
		mahasiswa = append(mahasiswa, m)
	}
	//fmt.Println("db nya:", mahasiswa)
	return mahasiswa, nil
}

func (rw *dbReadWriter) UpdateMahasiswa(mhs Mahasiswa) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateMahasiswa, mhs.NamaMahasiswa, mhs.Status, mhs.Nim)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadMahasiswaByNama(nama string) (Mahasiswa, error) {
	mahasiswa := Mahasiswa{NamaMahasiswa: nama}
	err := rw.db.QueryRow(selectMahasiswaByNama, nama).Scan(&mahasiswa.Nim, &mahasiswa.NamaMahasiswa,
		&mahasiswa.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Mahasiswa{}, err
	}

	return mahasiswa, nil
}
