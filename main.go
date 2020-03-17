package main

import (
	"UserService/apihandler"
	pb "UserService/proto"
	"fmt"
	codecs "github.com/amsokol/mongo-go-driver-protobuf"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)


const (
	defaultAtlasHost = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	//defaultAtlasHost = "mongodb+srv://nayan:tlwn722n@cluster0-8aov2.mongodb.net/test?retryWrites=true&w=majority"
	developmentMongoHost = "mongodb://dev-uni.cloudwalker.tv:6592"
	grpc_port        = ":7769"
	rest_port		 = ":7770"
)


func getMongoCollection(dbName, collectionName, mongoHost string )  *mongo.Collection {
	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err :=  mongo.Connect(ctx, options.Client().ApplyURI(mongoHost), options.Client().SetRegistry(reg))
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient.Database(dbName).Collection(collectionName)
}


func main() {
	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()


	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(rest_port, grpc_port)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()  // infinite loop
	log.Printf("Entering infinite loop")
	select {}
}


func startGRPCServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := apihandler.Server{
		getMongoCollection("cloudwalker", "users", defaultAtlasHost),
	}
	//serverOptions := []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamIntercept)}
	serverOptions := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)
	// attach the Ping service to the server
	pb.RegisterUserServiceServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(runtime.DefaultHeaderMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// Register ping
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}
