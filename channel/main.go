package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/nikogura/redisproxy/proxy/cache"
	"time"
)

type proxycache struct {
	c *cache.Cache
}

func (p proxycache) ServeHTTP(w http.ResponseWriter, r *http.Request) {


	fmt.Fprintf(w, "blah blah blah: %s", r.URL.Path)
}

func fetcher(key string, redisAddr string) (value interface{}, err error) {
	data := make(map[string]interface{})

	data["foo"] = "foo"
	data["bar"] = "bar"
	data["wip"] = "wip"
	data["zoz"] = "zoz"
	data["ten"] = 10

	if elem, ok := data[key]; ok {
		return elem, nil
	}

	return value, err
}


func main() {

	//send, receive := make(chan string), make(chan interface{})

	c := cache.NewCache(3, time.Second * 30, fetcher, time.Second * 5, "")

	err := http.ListenAndServe("0.0.0.0:8080", proxycache{c: c})
	log.Fatal(err)
}
