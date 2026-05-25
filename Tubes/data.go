package main

type Tagihan struct {
	Nama      string
	Kategori  string
	Nominal   float64
	JatuhTempo string
	Status    string
}

var daftarTagihan = []Tagihan{
	{"Listrik PLN", "Utilitas", 350000, "2026-05-15", "Belum"},
	{"Internet Indihome", "Utilitas", 450000, "2026-05-20", "Lunas"},
	{"Cicilan HP", "Elektronik", 800000, "2026-05-10", "Belum"},
	{"Netflix", "Hiburan", 54000, "2026-05-25", "Lunas"},
	{"BPJS Kesehatan", "Kesehatan", 150000, "2026-05-05", "Belum"},
	{"Spotify", "Hiburan", 54000, "2026-05-28", "Belum"},
	{"Air PDAM", "Utilitas", 120000, "2026-05-18", "Lunas"},
}
