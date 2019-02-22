package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc.test/protos/users"
	"log"
)

const ADDRESS = "127.0.0.1:1122"

func main() {
	clientConn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Printf("client dial %v err: %v", ADDRESS, err)
		return
	}

	client := pb.NewUsersClient(clientConn)
	req := new(pb.UserListsReq)
	page := new(pb.Page)
	page.PageNo = 1
	page.PageSize = 10
	page.Keyword = "keyword"
	req.Page = page

	ctx := context.Background()
	resp, err := client.Lists(ctx, req)
	if err != nil {
		log.Printf("client diag lists api err: %v", err)
		return
	}

	log.Printf("%s", resp)
}
