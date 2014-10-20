package flags

import (
	"fmt"
	"os"
	"strconv"

	"github.com/felixrabe/limited-netcat/errors"
)

func usage() {
	fmt.Println("Usage: netcat [-l] host port")
	os.Exit(1)
}

func Parse() (listen bool, host string, port int) {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		usage()
	}
	var err error
	if len(os.Args) == 3 {
		host = os.Args[1]
		if port, err = strconv.Atoi(os.Args[2]); err != nil {
			errors.Fatal(err)
		}
	} else {
		if os.Args[1] != "-l" {
			usage()
		}
		listen = true
		host = os.Args[2]
		if port, err = strconv.Atoi(os.Args[3]); err != nil {
			errors.Fatal(err)
		}
	}
	return
}
