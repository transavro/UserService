package apihandler

import (
	pb "UserService/proto"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"log"
	"time"
)

type Server struct {
	UserCollection *mongo.Collection
}

func (h *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.CreateReponse, error) {
	log.Printf("Create User Triggerd")
	// checking if the user is valid
	err := validatingData(req)
	if err != nil {
		return nil, makingError("Internal Server error", err.Error(), codes.Internal)
	}
	//inserting the user
	findResult := h.UserCollection.FindOne(ctx, bson.D{{"googleid", req.GetGoogleId()}})
	if findResult.Err() != nil {
		if findResult.Err() == mongo.ErrNoDocuments {
			ts, _ := ptypes.TimestampProto(time.Now())
			req.CreatedAt = ts
			req.LinkedDevices = []*pb.LinkedDevice{}
			_, err := h.UserCollection.InsertOne(context.Background(), req)
			if err != nil {
				return nil, makingError("Internal Server error", fmt.Sprintf("Error while creating user in Database. %s ", err.Error()), codes.Internal)
			}
			return &pb.CreateReponse{IsCreated: true}, nil
		} else {
			return nil, makingError("Internal Server error", fmt.Sprintf("Error while creating user in Database. %s ", err.Error()), codes.Internal)
		}
	} else {
		return nil, makingError("Internal Server error", "User Already Present", codes.Internal)
	}
}

func (h *Server) GetUser(ctx context.Context, request *pb.GetRequest) (*pb.User, error) {
	result := h.UserCollection.FindOne(context.Background(), bson.D{{"googleid", request.GetGoogleId()}})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, makingError("Internal Server error", fmt.Sprintf("User not found. %s \n", result.Err()), codes.NotFound)
		} else {
			return nil, makingError("Database Server error", fmt.Sprintf("Data base error. %s \n ", result.Err()), codes.Internal)
		}
	}
	var response *pb.User
	err := result.Decode(&response)
	if err != nil {
		return nil, makingError("Decoding", fmt.Sprintf("Error while decoding data %s \n", result.Err()), codes.Internal)
	}
	return response, err
}

func (h *Server) UpdateUser(ctx context.Context, req *pb.User) (*pb.UpdateResponse, error) {
	// check if the user data is valid
	err := validatingData(req)
	if err != nil {
		return nil, err
	}

	//replacing data
	_, err = h.UserCollection.ReplaceOne(ctx, bson.D{{"googleid", req.GetGoogleId()}}, &req)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.UpdateResponse{IsUpdated: false}, makingError("User not found.", "User data not present in Database.", codes.NotFound)
		} else {
			return &pb.UpdateResponse{IsUpdated: false}, makingError("Database error", fmt.Sprintf("Error while updating data %s \n", err.Error()), codes.Internal)
		}
	} else {
		return &pb.UpdateResponse{IsUpdated: true}, nil
	}
}

func (h *Server) DeleteUser(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteReponse, error) {
	log.Printf("Delete User hit")
	var response pb.DeleteReponse
	_, err := h.UserCollection.DeleteOne(ctx, bson.D{{"googleid", request.GetGoogleId()}})
	if err != nil {
		response.IsDeleted = false
	} else {
		response.IsDeleted = true
	}
	return &response, err
}

func (h *Server) LinkedTvDevice(ctx context.Context, request *pb.TvDevice) (*pb.LinkedDeviceResponse, error) {

	// if request is not proper
	if request.GetGoogleId() == "" {
		return nil, makingError("Request Error", fmt.Sprint("Google id cannot be empty"), codes.InvalidArgument)
	}else if request.GetLinkedDevice() == nil {
		return nil, makingError("Request Error", fmt.Sprintf("Linked Device cannot be empty for google id %s ", request.GetGoogleId()), codes.InvalidArgument)
	}

	var cwUser pb.User
	var response pb.LinkedDeviceResponse
	err := h.UserCollection.FindOne(ctx, bson.D{{"googleid", request.GetGoogleId()}}).Decode(&cwUser)
	if err != nil {
		return nil, makingError("Database error", fmt.Sprintf("User not found %s ", err.Error()), codes.NotFound)
	}
	// check if the device is already present
	for _, device := range cwUser.GetLinkedDevices() {
		if device.GetTvEmac() == request.GetLinkedDevice().GetTvEmac() {
			return nil, makingError("Database error", "Device already linked ", codes.AlreadyExists)
		}
	}
	cwUser.LinkedDevices = append(cwUser.LinkedDevices, request.LinkedDevice)
	_, err = h.UserCollection.ReplaceOne(ctx, bson.D{{"googleid", request.GetGoogleId()}}, cwUser)
	if err != nil {
		response.IsLinkedDeviceFetched = false
		return &response, makingError("Database error", fmt.Sprintf(err.Error()), codes.Internal)
	}
	response.LinkedDevices = cwUser.LinkedDevices
	response.IsLinkedDeviceFetched = true
	return &response, nil
}

func (h *Server) RemoveTvDevice(ctx context.Context, request *pb.RemoveTvDeviceRequest) (*pb.RemoveTvDeviceResponse, error) {
	var cwUser pb.User
	var response pb.RemoveTvDeviceResponse
	err := h.UserCollection.FindOne(context.Background(), bson.D{{"googleid", request.GoogleId}}).Decode(&cwUser)
	if err != nil {
		return nil, makingError("Database error", fmt.Sprintf("Error while fetching user data %s ", err.Error()), codes.Internal)
	}
	// Find and remove
	for i, v := range cwUser.LinkedDevices {
		if v.TvEmac == request.TvEmac {
			cwUser.LinkedDevices = append(cwUser.LinkedDevices[:i], cwUser.LinkedDevices[i+1:]...)
			_, err = h.UserCollection.ReplaceOne(ctx, bson.D{{"googleid", request.GoogleId}}, cwUser)
			if err != nil {
				return nil, makingError("Database error", fmt.Sprintf("Error while updating user data %s ", err.Error()), codes.Internal)
			}
			response.IsTvDeviceRemoved = true
			return &response, nil
		}
	}
	response.IsTvDeviceRemoved = false
	return nil, makingError("Database error", fmt.Sprintf("Tv device not found associated with user profile %s ", request.GetGoogleId()), codes.Internal)
}

func (h *Server) GetLinkedDevices(ctx context.Context, request *pb.GetRequest) (*pb.LinkedDeviceResponse, error) {
	var cwUser pb.User
	var response pb.LinkedDeviceResponse
	err := h.UserCollection.FindOne(ctx, bson.D{{"googleid", request.GoogleId}}).Decode(&cwUser)
	if err != nil {
		return nil, makingError("Database error", fmt.Sprintf("Error while fetching user data %s ", err.Error()), codes.Internal)
	}
	response.LinkedDevices = cwUser.LinkedDevices
	response.IsLinkedDeviceFetched = true
	return &response, nil
}

func (h *Server) GetLinkedUsers(ctx context.Context, linkUserReq *pb.LinkUserRequest) (*pb.LinkUserResponse, error) {

	if linkUserReq.GetEmac() == "" {
		return nil, makingError("Invalid Request", fmt.Sprintf("TV Emac cannot be empty"), codes.InvalidArgument)
	}

	// find all the user who have the requested tv emac
	result_cursor, err := h.UserCollection.Find(ctx, bson.M{"linkeddevices.tvemac": linkUserReq.GetEmac()})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, makingError("Database Error", fmt.Sprintf("No Users Linked to %s TV emac", linkUserReq.GetEmac()), codes.NotFound)
		}
		return nil, makingError("Database Error", fmt.Sprint(err.Error()), codes.Internal)
	}

	linkedUsers := new(pb.LinkUserResponse)
	for result_cursor.Next(ctx) {
		var user pb.User
		err = result_cursor.Decode(&user)
		if err != nil {
			return nil, makingError("Internal Error", fmt.Sprintf("Error while decoding data %s", err.Error()), codes.Internal)
		}
		linkedUsers.User = append(linkedUsers.User, &user)
	}
	return linkedUsers, nil
}
