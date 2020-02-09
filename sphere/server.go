package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/golang/protobuf/proto"
	spbuf "github.com/kantapapan/grpc-demo/sphere/protobuf"
)

var spList *spbuf.SphereList

func spheries(resp http.ResponseWriter, req *http.Request) {

	spData, err := proto.Marshal(spList)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/octet-stream")
	resp.Write(spData)
}

func main() {
	// load protobuf from csv
	list, err := createPbFromCsv("./sphere.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spList = list

	// start server
	http.HandleFunc("/", spheries)
	fmt.Println("Starting server on port 4041")
	if err := http.ListenAndServe(":4041", nil); err != nil {
		fmt.Println(err)
	}
}

// createPbFromCsv loads the currency data from csv into protobuf values
func createPbFromCsv(path string) (*spbuf.SphereList, error) {
	items := make([]*spbuf.Sphere, 0)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create CSV reader from file
	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		var num int32
		if i, err := strconv.Atoi(row[3]); err == nil {
			num = int32(i)
		}
		// create data row with protobuf
		c := &spbuf.Sphere{
			Job:     row[0],
			License: row[1],
			Skill:   row[2],
			Time:    num,
		}
		items = append(items, c)
	}
	return &spbuf.SphereList{Items: items}, err
}
