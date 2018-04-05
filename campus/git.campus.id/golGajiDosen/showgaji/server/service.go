package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Campus.mahasiswa.id"
	OnAdd     Status = 1
)

type GajiDosen struct {
	KdGGDosen  string
	JumlahGaji string
	Keterangan string
	Status     int32
	CreateBy   string
	CreateOn   string
	UpdateBy   string
	UpdateOn   string
}
type GajiDosens []GajiDosen

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
	AddGajiDosen(GajiDosen) error
	ReadGajiDosenByKdGGDosen(string) (GajiDosen, error)
	ReadGajiDosen() (GajiDosens, error)
	//UpdateMahasiswa(Mahasiswa) error
	//ReadMahasiswaByNama(string) (Mahasiswa, error)
}

type GajiDosenService interface {
	AddGajiDosenService(context.Context, GajiDosen) error
	ReadGajiDosenByKdGGDosenService(context.Context, string) (GajiDosen, error)
	ReadGajiDosenService(context.Context) (GajiDosens, error)
	//UpdateMahasiswaService(context.Context, Mahasiswa) error
	//ReadMahasiswaByNamaService(context.Context, string) (Mahasiswa, error)
}
