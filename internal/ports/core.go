package ports

// Files prot can list available files and find file string
type FilesPort interface {
	List() []string
	FindString(s string) (string, error)
}
