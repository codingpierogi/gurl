package print

import (
	"fmt"
	"os"
)

func Body(output string, content string) {
	if output == "" {
		fmt.Print(content)
		return
	}

	file, err := os.Create(output)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte(content))
}
