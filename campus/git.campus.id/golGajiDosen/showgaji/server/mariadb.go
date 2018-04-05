package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addGajiDosen               = `insert into golGajiDosen(Kd_GGDosen, Jumlah_Gaji, Keterangan, Status, CreateBy, CreateOn,UpdateBy, UpdateOn)values (?,?,?,?,?,?,?,?)`
	selectGajiDosenByKdGGDosen = `select Kd_GGDosen, Jumlah_Gaji, Keterangan, Status, CreateBy, CreateOn,UpdateBy, UpdateOn from golGajiDosen where Kd_GGDosen = ?`
	selectGajiDosen            = `select Kd_GGDosen, JUmlah_Gaji, Keterangan, Status, CreateBy, CreateOn,UpdateBy, UpdateOn from golGajiDosen where Status = '1'`
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

func (rw *dbReadWriter) AddGajiDosen(gajidosen GajiDosen) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addGajiDosen, gajidosen.KdGGDosen, gajidosen.JumlahGaji, gajidosen.Keterangan, gajidosen.Status,
		gajidosen.CreateBy, gajidosen.CreateOn, gajidosen.UpdateBy, gajidosen.UpdateOn, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadGajiDosenByKdGGDosen(gdsn string) (GajiDosen, error) {
	fmt.Println("show by gdsn")
	gajidosen := GajiDosen{KdGGDosen: gdsn}
	err := rw.db.QueryRow(selectGajiDosenByKdGGDosen, gdsn).Scan(&gajidosen.KdGGDosen, &gajidosen.JumlahGaji, &gajidosen.Keterangan,
		&gajidosen.Status, &gajidosen.CreateBy, &gajidosen.CreateOn, &gajidosen.UpdateBy, &gajidosen.UpdateOn)

	if err != nil {
		return GajiDosen{}, err
	}

	return gajidosen, nil
}

func (rw *dbReadWriter) ReadGajiDosen() (GajiDosens, error) {
	fmt.Println("show all")
	gajidosen := GajiDosens{}
	rows, _ := rw.db.Query(selectGajiDosen)
	defer rows.Close()
	for rows.Next() {
		var m GajiDosen
		err := rows.Scan(&m.KdGGDosen, &m.JumlahGaji, &m.Keterangan, &m.Status, &m.CreateBy, &m.CreateOn, &m.UpdateBy, &m.UpdateOn)
		if err != nil {
			fmt.Println("error query:", err)
			return gajidosen, err
		}
		gajidosen = append(gajidosen, m)
	}
	//fmt.Println("db nya:", mahasiswa)
	return gajidosen, nil
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
