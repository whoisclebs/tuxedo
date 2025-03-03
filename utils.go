package tuxedo

import (
	"io"
	"log"
)

func doClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
