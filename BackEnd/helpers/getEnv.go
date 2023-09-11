package helpers

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, bool) {
	val, ok := os.LookupEnv(key)
	fmt.Println("This is secret " + key + " value: " + val)
	return val, ok
}
