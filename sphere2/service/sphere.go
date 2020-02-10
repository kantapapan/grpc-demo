package service

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	pb "github.com/kantapapan/grpc-demo/sphere2/pb"
)

// SphereService ...
type SphereService struct {
}

var spList *pb.SphereList

// GetSphereByJob ...
func (s *SphereService) GetSphereByJob(
	ctx context.Context,
	message *pb.SphereJobMessage) (*pb.SphereList, error) {
	spList, err := createPbFromCsv("./sphere.csv")

	if err != nil {
		fmt.Printf("%+v", err)
	}

	/*
		fmt.Println("----")
		fmt.Printf("%+v", spList)
		fmt.Println("----")
	*/

	switch message.TargetJob {
	case "Architect":
		return spList, nil
	case "CloudEngineer":
		return spList, nil
	}

	return nil, errors.New("Not Found YourSpere")
}

// createPbFromCsv loads the currency data from csv into protobuf values
func createPbFromCsv(path string) (*pb.SphereList, error) {
	items := make([]*pb.Sphere, 0)
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

		var id int32
		if i, err := strconv.Atoi(row[0]); err == nil {
			id = int32(i)
		}

		var num int32
		if i, err := strconv.Atoi(row[3]); err == nil {
			num = int32(i)
		}
		// create data row with protobuf
		c := &pb.Sphere{
			Id:      id,
			Job:     row[1],
			License: row[2],
			Skill:   row[3],
			Time:    num,
		}
		items = append(items, c)
	}
	return &pb.SphereList{Items: items}, err
}
