package main

import (
	"fmt"
	"time"

	"github.com/sfi2k7/bluecache"
)

func main() {
	// c := bluecache.New()

	// start := time.Now()
	// for x := 0; x < 10000; x++ {
	// 	c.Set([]byte("Key"+strconv.Itoa(x)), []byte("Value"+strconv.Itoa(x)))
	// }
	// fmt.Println(time.Since(start))
	// start = time.Now()

	// var v []byte
	// var err error
	// for x := 0; x < 10000; x++ {
	// 	v, err = c.Get([]byte("Key" + strconv.Itoa(x)))
	// }

	// fmt.Println(time.Since(start))
	// fmt.Println(string(v), err)
	bc := bluecache.NewWithBackend(bluecache.BackendBolt)
	bc.Set([]byte("Key"), []byte("Value"))
	start := time.Now()
	for x := 0; x < 100000; x++ {
		bc.Get([]byte("Key"))
	}
	fmt.Println(time.Since(start))
}

/*
package main

//"gopkg.in/redis.v3"

// type Person struct {
// 	MD MyDate
// }

// type MyDate time.Time

// func (md *MyDate) MarshalJson() ([]byte, error) {
// 	return []byte("\"Data\""), nil
// }

// func main() {
// 	// p :=&Person{MD:MyDate(time.Now())}
// 	// jsoned,err:=json.Marshal(p)
// 	// fmt.Println(err)
// 	// fmt.Println(string(jsoned))

// 	fmt.Println(os.Getwd())
// }

// func connect() *redis.Client {
// 	c := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Network: "tcp"})
// 	return c
// }

// func main()  {
//      c:=connect()
//      defer c.Close()
//      c.Auth("passme")
//      fmt.Println(c.Ping())
//      c.Set("hello", "world", 0)
//      //r,_:= sm.Result()
//     //  fmt.Println(sm.Err() )
//     //  fmt.Println(sm.Result())
//     //  fmt.Println(sm.String())
//     //  fmt.Println(sm.Val())
//     sm:=c.Get("hello2")
//     bts,err:= sm.Bytes()
//     fmt.Println(bts,err)

//     // _,err:= sm.Bytes()
//     // fmt.Println(sm.Err() )
//     // fmt.Println(err.Error() )
// }

*/
