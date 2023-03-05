package ports

// Files prot can list available files and find file string
type FilesPort interface {
	// List implements the behavior of packr2 which will return a slice of
	// strings which are the file paths
	List() []string
	// FindString implements the behavior of packr2 which will take a string
	// path to a file in the Box and return the contents
	FindString(s string) (string, error)
}
