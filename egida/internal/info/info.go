package info

import (
	"context"
	"errors"
	"fmt"
	"github.com/antonioalfa22/go-utils/collections"
	"time"

	grpc "github.com/antonioalfa22/egida/proto"
)

type Result struct {
	Host string
	Lines []string
}

func GetWorkerInfo(hostslist []string, services string, packages string, hardening string) {
	if services != "" {
		if services == "stopped" {
			all, err :=GetStoppedServices(hostslist)
			if err != nil {
				fmt.Println(err)
				return
			}
			collections.ForEach(all, func(result Result) {
				fmt.Println("-------> Host", result.Host)
				for _, line := range result.Lines {fmt.Println(line)}
			})
		} else if services == "running" {
			all, err :=GetRunningServices(hostslist)
			if err != nil {
				fmt.Println(err)
				return
			}
			collections.ForEach(all, func(result Result) {
				fmt.Println("-------> Host", result.Host)
				for _, line := range result.Lines {fmt.Println(line)}
			})
		} else {
			all, err := GetAllServices(hostslist)
			if err != nil {
				fmt.Println(err)
				return
			}
			collections.ForEach(all, func(result Result) {
				fmt.Println("-------> Host", result.Host)
				for _, line := range result.Lines {fmt.Println(line)}
			})
		}
	}
	if packages != "" {
		if packages == "all" {
			all, err :=GetAllPackages(hostslist)
			if err != nil {
				fmt.Println(err)
				return
			}
			collections.ForEach(all, func(result Result) {
				fmt.Println("-------> Host", result.Host)
				for _, line := range result.Lines {fmt.Println(line)}
			})
		}
	}
	if hardening != "" {
		if hardening == "lynis" {
			all, err :=GetLynisScore(hostslist)
			if err != nil {
				fmt.Println(err)
				return
			}
			collections.ForEach(all, func(result Result) {
				fmt.Println("-------> Host", result.Host)
				for _, line := range result.Lines {fmt.Println(line)}
			})
		}
	}
}

func GetAllServices(hosts []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var results []Result
	for _, h := range hosts {
		r := Result{Host: h, Lines: []string{}}
		client := CreateServicesClient(h)
		res, err := client.ListAllServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			for _, s := range res.Services {
				r.Lines = append(r.Lines, s.Name+"->"+s.Status)
			}
			results = append(results, r)
		}
	}
	return results, nil
}

func GetRunningServices(hosts []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var results []Result
	for _, h := range hosts {
		r := Result{Host: h, Lines: []string{}}
		client := CreateServicesClient(h)
		res, err := client.ListRunningServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			for _, s := range res.Services {
				r.Lines = append(r.Lines, s.Name+"->"+s.Status)
			}
			results = append(results, r)
		}
	}
	return results, nil
}

func GetStoppedServices(hosts []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var results []Result
	for _, h := range hosts {
		r := Result{Host: h, Lines: []string{}}
		client := CreateServicesClient(h)
		res, err := client.ListStoppedServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			for _, s := range res.Services {
				r.Lines = append(r.Lines, s.Name+"->"+s.Status)
			}
			results = append(results, r)
		}
	}
	return results, nil
}

func GetAllPackages(hosts []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var results []Result
	for _, h := range hosts {
		r := Result{Host: h, Lines: []string{}}
		client := CreatePackagesClient(h)
		res, err := client.ListAllPackages(ctx, &grpc.ListPackagesReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			for _, s := range res.Packages {
				r.Lines = append(r.Lines, s.Name)
			}
			results = append(results, r)
		}
	}
	return results, nil
}

func GetLynisScore(hosts []string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	var results []Result
	for _, h := range hosts {
		r := Result{Host: h, Lines: []string{}}
		client := CreateHardeningClient(h)
		res, err := client.GetLynisScore(ctx, &grpc.ScoreReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			r.Lines = res.Log
			results = append(results, r)
		}
	}
	return results, nil
}

func GetMachineHasGUI(hosts []string) ([]bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	var results []bool
	for _, h := range hosts {
		client := CreateMachineInfoClient(h)
		res, err := client.MachineHasGUI(ctx, &grpc.HasGUIReq{})
		if err != nil {
			return nil, errors.New("Error: can not connect to host " + h + " on port 8128")
		} else {
			results = append(results, res.Gui)
		}
	}
	return results, nil
}

func GetMachineOpenPorts(host string) ([]int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	var results []int64
	client := CreateMachineInfoClient(host)
	res, err := client.ListOpenPorts(ctx, &grpc.ListOpenPortsReq{})
	if err != nil {
		return nil, errors.New("Error: can not connect to host " + host + " on port 8128")
	} else {
		for _, port := range res.Ports {
			results = append(results, port.Number)
		}
	}
	return results, nil
}