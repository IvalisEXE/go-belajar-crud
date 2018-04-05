package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDosen                = `insert into Dosen(Kd_Dosen, Nama_Dosen, Status, CreateBy, CreateOn)values (?,?,?,?,?)`
	selectDosenByKdDosen    = `select Kd_Dosen, Nama_Dosen, Status, CreateBy from Dosen where Kd_Dosen = ?`
	selectDosenByKeterangan = `select Keterangan, Nama_Dosen, Status, CreateBy from Dosen where Keterangan = ?`
	selectDosen             = `select Kd_Dosen, Nama_Dosen, Status, CreateBy from Dosen where Status = '1'`
	//updateMahasiswa       = `update Mahasiswa set Nim=?, Nama_Mahasiswa=?, Status=? where Nim=?`
	//selectMahasiswaByNama = `select Nim,Nama_Mahasiswa, Status from Mahasiswa where Nama_Mahasiswa=?`
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
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDosen, dosen.KdDosen, dosen.NamaDosen, OnAdd, dosen.CreateBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadDosenByKdDosen(kddosen string) (Dosen, error) {
	fmt.Println("show by kddosen")
	dosen := Dosen{KdDosen: kddosen}
	err := rw.db.QueryRow(selectDosenByKdDosen, kddosen).Scan(&dosen.KdDosen, &dosen.NamaDosen, &dosen.Status, &dosen.CreateBy)

	if err != nil {
		return Dosen{}, err
	}

	return dosen, nil
}

func (rw *dbReadWriter) ReadDosenByKeterangan(ktrg string) (Dosen, error) {
	fmt.Println("show by ktrg")
	dosen := Dosen{Keterangan: ktrg}
	err := rw.db.QueryRow(selectDosenByKeterangan, ktrg).Scan(&dosen.Keterangan, &dosen.NamaDosen, &dosen.Status, &dosen.CreateBy)

	if err != nil {
		return Dosen{}, err
	}

	return dosen, nil
}

func (rw *dbReadWriter) ReadDosen() (Dosens, error) {
	fmt.Println("show all")
	dosen := Dosens{}
	rows, _ := rw.db.Query(selectDosen)
	defer rows.Close()
	for rows.Next() {
		var m Dosen
		err := rows.Scan(&m.KdDosen, &m.NamaDosen, &m.Status, &m.CreateBy)
		if err != nil {
			fmt.Println("error query:", err)
			return dosen, err
		}
		dosen = append(dosen, m)
	}
	//fmt.Println("db nya:", mahasiswa)
	return dosen, nil
}

/*
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
*/
