package info

import (
	"fmt"
	grpc "github.com/antonioalfa22/egida/proto"
	googlegrpc "google.golang.org/grpc"
	"log"
)

func CreateServicesClient(host string) grpc.ServicesClient {
	addr := fmt.Sprintf("%s:%s", host, "8128")
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("Impossible connect: %v", err)
	}
	return grpc.NewServicesClient(conn)
}

func CreatePackagesClient(host string) grpc.PackagesClient {
	addr := fmt.Sprintf("%s:%s", host, "8128")
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("Impossible connect: %v", err)
	}
	return grpc.NewPackagesClient(conn)
}

func CreateHardeningClient(host string) grpc.HardeningScoresClient {
	addr := fmt.Sprintf("%s:%s", host, "8128")
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("Impossible connect: %v", err)
	}
	return grpc.NewHardeningScoresClient(conn)
}

func CreateMachineInfoClient(host string) grpc.MachineInfoClient {
	addr := fmt.Sprintf("%s:%s", host, "8128")
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("Impossible connect: %v", err)
	}
	return grpc.NewMachineInfoClient(conn)
}