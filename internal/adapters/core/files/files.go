package files

// 1. Create struct

type Adapter struct {
}

// 3. Build factory function

func NewAdapter() *Adapter {
	return &Adapter{}
}

// 2. Build methods on interface

func (a Adapter) List() []string {
	s := make([]string, 1)
	return s
}

func (a Adapter) FindString(s string) (string, error) {
	return "file_name", nil
}
