package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v4"
	pb "github.com/tech-with-moss/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

<<<<<<< HEAD
func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

type UserManagementServer struct {
	conn                *pgx.Conn
	first_user_creation bool
	pb.UnimplementedUserManagementServer
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	return s.Serve(lis)
}

// When user is added, read full userlist from file into
// userlist struct, then append new user and write new userlist back to file
func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {

	createSql := `
	CREATE TABLE IF NOT EXIST roll_assignment (
		assignment_id INTEGER,
		assignee VARCHAR(255) NOT NULL,
		role_id INTEGER NOT NULL,
		status INTEGER,
		tenant_id INTEGER,
		service_id INTEGER,
		is_group INTEGER,
	  );
	`
	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	server.first_user_creation = false

	log.Printf("Received: %v %v %v %v %v %v %v", in.GetRollAssignmentId(), in.GetAssaignee(), in.GetRollId(), in.GetStatus(), in.GetTenantId(), in.GetServiceId(), in.GetIsGroup())

	created_user := &pb.User{RollAssignmentId: in.GetRollAssignmentId(), Assaignee: in.GetAssaignee(), RollId: in.GetRollId(), Status: in.GetStatus(), TenantId: in.GetTenantId(), ServiceId: in.GetServiceId(), IsGroup: in.GetIsGroup()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into roll_assignment(assignment, assaignee, roll_id, status, tenant_id, service_id) values ($1,$2,$3,$4,$5,$6,$7)",
		created_user.RollAssignmentId, created_user.Assaignee, created_user.RollId, created_user.Status, created_user.TenantId, created_user.ServiceId, created_user.IsGroup)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return created_user, nil

}

func (server *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UsersList, error) {

	var users_list *pb.UsersList = &pb.UsersList{}
	rows, err := server.conn.Query(context.Background(), "select * from roll_assignment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.IsGroup, &user.ServiceId, &user.TenantId, &user.Status, &user.RollId, &user.Assaignee, &user.RollAssignmentId)
		if err != nil {
			return nil, err
		}
		users_list.Users = append(users_list.Users, &user)

	}
	return users_list, nil
}

func main() {
	database_url := "postgres://postgres:sanjai@localhost:5432/postgres"
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	user_mgmt_server.conn = conn
	user_mgmt_server.first_user_creation = true
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
=======
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
>>>>>>> origin/main
}
