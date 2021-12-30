package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	// change program working directory
	err := os.Chdir("/")
	if err != nil {
		fmt.Println(err)
		return
	}
	user := os.Getenv("USER")

	// get input from flags
	saveAsFileType := flag.String("t", "text", "type what kind of files type you want to convert text/json")
	// flag.Parse()
	tipe := *saveAsFileType
	outputFileLocation := flag.String("o", fmt.Sprintf("/home/%v/Downloads/file.%v", user, tipe), "type where you want to save the file")
	flag.Parse()

	// get input from terminal argument
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("please input a file location in terminal argument")
		return
	}
	fileLocation := args[0]

	switch *saveAsFileType {
	case "json":
		file, _ := readfile(fileLocation)
		fileJson, _ := ConvertToJSON(file)
		err = WriteFile(*outputFileLocation, fileJson)
		if err != nil {
			fmt.Println(err)
		}
	case "text":
		file, _ := readfile(fileLocation)
		err = WriteFile(*outputFileLocation, file)
		if err != nil {
			fmt.Println(err)
		}
	default:
		file, _ := readfile(fileLocation)
		err = WriteFile(*outputFileLocation, file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func readfile(filepath string) (string, error) {
	// read file
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(file), nil
}

func ConvertToJSON(file string) (string, error) {
	// convert to json
	fileJson, err := json.Marshal(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(fileJson), nil
}

func WriteFile(path string, content string) error {

	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return err
		}
		defer file.Close()
	}

	// buka file dengan level akses READ & WRITE
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return err
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString(content)
	if isError(err) {
		return err
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return err
	}

	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
