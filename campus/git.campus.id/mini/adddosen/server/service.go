package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Campus.dosen.id"
	OnAdd     Status = 1
)

type Dosen struct {
	KdDosen   string
	NamaDosen string
	Keterangan string
	Status    int32
	CreateBy  string
}
type Dosens []Dosen

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddDosen(Dosen) error
	ReadDosenByKdDosen(int32) (Mahasiswa, error)
	ReadDosen() (Mahasiswas, error)
	UpdateDosen(Mahasiswa) error
	ReadDosenByKeterangan(string) (Mahasiswa, error)
}

type DosenService interface {
	AddDosenService(context.Context, Dosen) error
	ReadDosenByKdDosenService(context.Context, int32) (Dosen, error)
	ReadDosenService(context.Context) (Dosens, error)
	UpdateDosenService(context.Context, Dosen) error
	ReadDosenByKeteranganService(context.Context, string) (Dosen, error)
}
