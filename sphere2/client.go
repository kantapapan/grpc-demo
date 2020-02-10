package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kantapapan/grpc-demo/sphere2/pb"

	"google.golang.org/grpc"
)

// main ...
func main() {

	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewSphereServiceClient(conn)
	//message := &pb.SphereJobMessage{TargetJob: "Architect"}
	message := &pb.SphereJobMessage{TargetJob: "CloudEngineer"}

	res, err := client.GetSphereByJob(context.TODO(), message)
	fmt.Printf("result:%+v \n", res)
	fmt.Printf("error::%+v \n", err)
}
