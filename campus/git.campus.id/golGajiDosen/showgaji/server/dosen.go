package server

import (
	"context"
)

type gajidosen struct {
	writer ReadWriter
}

func NewGajiDosen(writer ReadWriter) GajiDosenService {
	return &gajidosen{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *gajidosen) AddGajiDosenService(ctx context.Context, gajidosen GajiDosen) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddGajiDosen(gajidosen)
	if err != nil {
		return err
	}

	return nil
}

func (c *gajidosen) ReadGajiDosenByKdGGDosenService(ctx context.Context, mob string) (GajiDosen, error) {
	gd, err := c.writer.ReadGajiDosenByKdGGDosen(mob)
	//fmt.Println(mhs)
	if err != nil {
		return gd, err
	}
	return gd, nil
}

func (c *gajidosen) ReadGajiDosenService(ctx context.Context) (GajiDosens, error) {
	gd, err := c.writer.ReadGajiDosen()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return gd, err
	}
	return gd, nil
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
