package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

import pb "grpc.test/protos/users"

const ADDRESS = "127.0.0.1:1122"

type userService struct {
}

func (u userService) Lists(ctx context.Context, req *pb.UserListsReq) (*pb.UserListsResp, error) {
	log.Printf("server receive msg %s", req)
	resp := new(pb.UserListsResp)
	addrs1 := []*pb.Addr{
		&pb.Addr{Id: 12, Name: "test1", Mobile: "156232323"},
		&pb.Addr{Id: 13, Name: "test2", Mobile: "156232324"},
	}
	addrs2 := []*pb.Addr{
		&pb.Addr{Id: 18, Name: "test18", Mobile: "156232328"},
		&pb.Addr{Id: 19, Name: "test19", Mobile: "156232329"},
	}
	users := []*pb.User{
		&pb.User{Id: 1, Name: "qiang1", Addresses: addrs1},
		&pb.User{Id: 2, Name: "qiang2", Addresses: addrs2},
	}
	resp.Users = users

	return resp, nil
}

var userS = userService{}

func main() {
	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Printf("server listen %v err: %v", ADDRESS, err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterUsersServer(s, userS)

	log.Printf("server start $v ... ", ADDRESS);
	s.Serve(listener)
}
