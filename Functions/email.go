package functions

import (
	"fmt"
	"os"
)

func Email() {

	from := os.Getenv("emailAddr")

	//password := os.Getenv("emailPass")

	fmt.Println(from)
}
