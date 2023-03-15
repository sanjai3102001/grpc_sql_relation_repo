package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// var user_id = "sanjai"
	// var new_users1 = make(map[string]string)
	// new_users1["user_id"] = "sanjai"
	// var arr [5]int32
	// myarr := [5]int{1, 2, 3, 4, 5}
	roll_id := 2
	status := 1
	tenant_id := 3
	is_group := 1
	service_id := 2
	var new_users = make(map[int]string)
	new_users[5] = "Sanjai"
	new_users[4] = "Dhanalakshmi"
	new_users[3] = "saravana"
	new_users[2] = "Adish"
	new_users[1] = "Haris"
	// for service_id, name, owner_id, prefix, action := range new_users {
	// 	r, err := c.CreateNewUser(ctx, &pb.NewUser{ServiceId: int32(service_id), Name: string(name), OwnerId: int32(owner_id), Prefix: int32(prefix), Action: int32(action)})
	for assignment_id, assaignee := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{AssignmentId: int32(assignment_id), Assaignee: string(assaignee), RollId: int32(roll_id), Status: int32(status), TenantId: int32(tenant_id), ServiceId: int32(service_id), IsGroup: int32(is_group)})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`service Details:
assaignment_id: %d
assaignee: %s
status: %d
tenant_id: %d
service_id: %d
is_group: %d`, r.GetAssignmentId(), r.GetAssaignee(), r.GetRollId(), r.GetStatus(), r.GetTenantId(), r.GetServiceId(), r.GetIsGroup())

	}
	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not create roll_assaignment: %v", err)
	}
	log.Print("\nUSER LIST: \n")
	fmt.Printf("r.GetRoles(): %v\n", r.GetUsers())
}
