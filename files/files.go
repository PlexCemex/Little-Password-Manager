package files

import (
	"fmt"
	"os"
	"test/test_app_4/output"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
	}
	fmt.Println("Successfully")
}
func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
