package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Open() {
	f, err := os.Open("test.cfg")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
