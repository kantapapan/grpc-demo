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

	filterByJob("/", spList)

	if err != nil {
		fmt.Printf("%+v", err)
	}

	switch message.TargetJob {
	case "Architect":
		return spList, nil
	case "CloudEngineer":
		return spList, nil
	}

	return nil, errors.New("Not Found YourSpere")
}

//  filterByJob ...
func filterByJob(job string, l *pb.SphereList) (*pb.SphereList, error) {
	for _, item := range l.Items {
		fmt.Printf("%+v\n", item)
	}
	return spList, nil
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

		var job string
		job = row[1]

		var num int32
		if i, err := strconv.Atoi(row[3]); err == nil {
			num = int32(i)
		}
		// create data row with protobuf
		c := &pb.Sphere{
			Id:      id,
			Job:     job,
			License: row[2],
			Skill:   row[3],
			Time:    num,
		}
		items = append(items, c)
	}
	return &pb.SphereList{Items: items}, err
}
