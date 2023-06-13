package helpers

import (
	"fmt"
	"os"
)

func GetEnv(key string) {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf
	}
}
