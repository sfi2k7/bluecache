package main

import (
	"fmt"

	"github.com/sfi2k7/bluecache"

	"time"

	"strconv"
)

func main() {
	c := bluecache.New()

	start := time.Now()
	for x := 0; x < 10000; x++ {
		c.Set([]byte("Key"+strconv.Itoa(x)), []byte("Value"+strconv.Itoa(x)))
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	var v []byte
	var err error
	for x := 0; x < 10000; x++ {
		v, err = c.Get([]byte("Key" + strconv.Itoa(x)))
	}
	fmt.Println(time.Since(start))
	fmt.Println(string(v), err)
}
