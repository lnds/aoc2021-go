package shared

import (
	"log"
	"strconv"
)

func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
