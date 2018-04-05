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
	//fmt.Println("dosen")
	err := c.writer.AddDosen(dosen)
	if err != nil {
		return err
	}

	return nil
}

/*
func (c *mahasiswa) ReadMahasiswaByNimService(ctx context.Context, mob int32) (Mahasiswa, error) {
	mhs, err := c.writer.ReadMahasiswaByNim(mob)
	//fmt.Println(mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}

func (c *mahasiswa) ReadMahasiswaService(ctx context.Context) (Mahasiswas, error) {
	mhs, err := c.writer.ReadMahasiswa()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return mhs, err
	}
	return mhs, nil
}

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
