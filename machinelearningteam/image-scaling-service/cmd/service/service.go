package main

import (
	"log"
	"net"
	"net/http"

	"github.com/e-conomic/hiring-assigments/machinelearningteam/image-scaling-service/pkg/api"
	pb "github.com/e-conomic/hiring-assigments/machinelearningteam/image-scaling-service/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// starts the Prometheus stats endpoint server
func startPromHTTPServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("prometheus err", port)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go startPromHTTPServer("5001")

	s := grpc.NewServer()
	pb.RegisterImageScalerServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
