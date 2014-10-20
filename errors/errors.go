package errors

import (
	"fmt"
	"os"
)

// Fatal outputs errors just like log.Fatal() would, but without a timestamp.
func Fatal(e error) {
	fmt.Println(e)
	os.Exit(1)
}
