package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, fmt.Errorf("File not Exsist")
		}

		return false, err
	}

	return true, nil
}
func CreateFile(path string, data string) error {
	isExist, err := FileExists(path)

	if !isExist {
		return err
	}

	file, err := os.Create(path)

	if err != nil {
		return fmt.Errorf("fail to creat file")
	}

	defer file.Close()

	n, err := file.WriteString(data)

	if err != nil {
		return fmt.Errorf("fail to write data")
	}

	fmt.Print("create file and store data total lenght is ", n)

	return nil
}

func ReadFileWithoutOpen(path string) ([]byte, error) {
	rawByte, err := os.ReadFile(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return rawByte, fmt.Errorf("No File Found")
		} else {
			return rawByte, err
		}
	}
	return rawByte, nil
}

func ReadFileWithOpen(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("No File Found")
		}
		return nil, err
	}

	defer file.Close()

	return nil, nil

}

func main() {
	filePath := filepath.Join("Book", "data.txt")
	if err := CreateFile(filePath, "THis is test data"); err != nil {
		fmt.Println(err)
	}

	byteData, err := ReadFileWithoutOpen(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(byteData))
}
