package doczip

import (
	"fmt"
	"os"
)

func main() {
	zipfile, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
	}

	defer zipfile.Close()
}
