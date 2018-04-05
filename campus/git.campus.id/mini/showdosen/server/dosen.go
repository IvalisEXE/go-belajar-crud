package server

import (
	"context"
)

type dosen struct {
	writer ReadWriter
}

func NewDosen(writer ReadWriter) DosenService {
	return &dosen{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *dosen) AddDosenService(ctx context.Context, dosen Dosen) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddDosen(dosen)
	if err != nil {
		return err
	}

	return nil
}

func (c *dosen) ReadDosenByKdDosenService(ctx context.Context, mob string) (Dosen, error) {
	dsn, err := c.writer.ReadDosenByKdDosen(mob)
	//fmt.Println(mhs)
	if err != nil {
		return dsn, err
	}
	return dsn, nil
}

func (c *dosen) ReadDosenByKeteranganService(ctx context.Context, mob string) (Dosen, error) {
	dsn, err := c.writer.ReadDosenByKeterangan(mob)
	//fmt.Println(mhs)
	if err != nil {
		return dsn, err
	}
	return dsn, nil
}

func (c *dosen) ReadDosenService(ctx context.Context) (Dosens, error) {
	dsn, err := c.writer.ReadDosen()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return dsn, err
	}
	return dsn, nil
}

/*
func (c *mahasiswa) UpdateMahasiswaService(ctx context.Context, mhs Mahasiswa) error {
	err := c.writer.UpdateMahasiswa(mhs)
	if err != nil {
		return err
	}
	return nil
}

func (c *mahasiswa) ReadMahasiswaByNamaService(ctx context.Context, nama string) (Mahasiswa, error) {
	mhs, err := c.writer.ReadMahasiswaByNama(nama)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}
*/
