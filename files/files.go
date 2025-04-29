package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	var data, err = os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content string, filename string) {
	var file, err = os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File written successfully")
}
