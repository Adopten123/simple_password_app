package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content string, filename string) {
	var file, err = os.Create(filename)

	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File written successfully")
	}

	file.Close()
}
