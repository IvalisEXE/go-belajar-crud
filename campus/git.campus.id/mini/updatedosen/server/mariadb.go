package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDosen             = `insert into Dosen(Kd_Dosen, Nama_Dosen, Status, CreateBy, CreateOn)values (?,?,?,?,?)`
	selectDosenByKdDosen = `select Kd_Dosen, Nama_Dosen, Status, CreateBy from Dosen where Kd_Dosen = ?`
	selectDosen          = `select Kd_Dosen, Nama_Dosen, Status, CreateBy from Dosen Where Status ='1'`
	updateDosen          = `update Dosen set Kd_Dosen=?, Nama_Dosen=?, Status=?, UpdateBy=?, UpdateOn=? where Kd_Dosen=?`
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

func (rw *dbReadWriter) UpdateDosen(dsn Dosen) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateDosen, dsn.KdDosen, dsn.NamaDosen, dsn.Status, dsn.UpdateBy, time.Now(), dsn.KdDosen)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

/*
func (rw *dbReadWriter) ReadMahasiswaByNama(nama string) (Mahasiswa, error) {
	mahasiswa := Mahasiswa{NamaMahasiswa: nama}
	err := rw.db.QueryRow(selectMahasiswaByNama, nama).Scan(&mahasiswa.Nim, &mahasiswa.NamaMahasiswa,
		&mahasiswa.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Mahasiswa{}, err
	}

	return mahasiswa, nil
}*/
