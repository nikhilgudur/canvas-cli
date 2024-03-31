package errors

import (
	"log"
)

func ErrorLog(err error) {
	log.Fatalf("%v", err)
}
