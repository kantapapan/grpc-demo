package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	spbuf "github.com/kantapapan/grpc-demo/sphere/protobuf"
)

func main() {
	resp, err := http.Get("http://localhost:4041/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// read data from body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal
	spList := new(spbuf.SphereList)
	if err := proto.Unmarshal(data, spList); err != nil {
		log.Fatalln(err)
	}

	// display protobuf data
	for _, item := range spList.Items {
		fmt.Printf("%+v\n", item)
	}
}
