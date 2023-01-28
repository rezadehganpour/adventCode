package model

type Directory struct {
	Name        string
	FullName    string
	Files       []File
	TotalSize   int
	Directories []*Directory
	Parent      *Directory
}
