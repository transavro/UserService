package apihandler

import (
	pb "UserService/proto"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Server struct {
	*mongo.Collection
}

func (h *Server) AddUser(ctx context.Context, req *pb.User) (*empty.Empty, error) {
	findResult := h.FindOne(ctx, bson.D{{"googleid", req.GetGoogleId()}})
	if findResult.Err() != nil {
		if findResult.Err() == mongo.ErrNoDocuments {
			ts, _ := ptypes.TimestampProto(time.Now())
			req.CreatedAt = ts
			req.LinkedDevices = []*pb.LinkedDevice{}
			_, err := h.InsertOne(context.Background(), req)
			if err != nil {
				return nil, makingError("Internal Server error", fmt.Sprintf("Error while creating user in Database. %s ", err.Error()), codes.Internal)
			}
			return new(empty.Empty), nil
		} else {
			return nil, makingError("Internal Server error", fmt.Sprintf("Error while creating user in Database. %s ", findResult.Err()), codes.Internal)
		}
	} else {
		return nil, makingError("Internal Server error", "User Already Present", codes.Internal)
	}
}

func (h *Server) ListUsers(_ *empty.Empty, stream pb.UserService_ListUsersServer) error {
	cur, err := h.Find(stream.Context(), bson.D{})
	if err != nil {
		return err
	}
	defer cur.Close(stream.Context())
	defer stream.Context().Done()
	user := new(pb.User)
	for cur.Next(stream.Context()) {
		if err = cur.Decode(user); err != nil {
			return err
		}
		if err = stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (h *Server) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	//replacing data
	_, err := h.ReplaceOne(ctx, bson.D{{"googleid", req.GetGoogleId()}}, &req)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, makingError("User not found.", "User data not present in Database.", codes.NotFound)
		} else {
			return nil, makingError("Database error", fmt.Sprintf("Error while updating data %s \n", err.Error()), codes.Internal)
		}
	} else {
		return req, nil
	}
}

func (h *Server) GetUser(ctx context.Context, req *pb.GetRequest) (*pb.User, error) {
	result := h.FindOne(ctx, bson.M{"googleid": req.GetGoogleId()})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, makingError("Internal Server error", fmt.Sprintf("User not found. %s \n", result.Err()), codes.NotFound)
		} else {
			return nil, makingError("Database Server error", fmt.Sprintf("Data base error. %s \n ", result.Err()), codes.Internal)
		}
	}
	var resp pb.User
	if err := result.Decode(&resp); err != nil {
		return nil, makingError("Decoding", fmt.Sprintf("Error while decoding data %s \n", result.Err()), codes.Internal)
	}
	return &resp, nil
}

func (h *Server) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReponse, error) {
	if _, err := h.DeleteOne(ctx, bson.M{"googleid": req.GetGoogleId()}); err != nil {
		return nil, makingError("Database Server error", fmt.Sprintf("Data base error. %s \n ", err.Error()), codes.Internal)
	}
	return &pb.DeleteReponse{IsDeleted: true}, nil
}

func (h *Server) AddTvDevice(ctx context.Context, request *pb.TvDevice) (*pb.LinkedDeviceResponse, error) {
	var cwUser pb.User
	var response pb.LinkedDeviceResponse
	var err error
	if err = h.FindOne(ctx, bson.M{"googleid": request.GetGoogleId()}).Decode(&cwUser); err != nil {
		return nil, makingError("Database error", fmt.Sprintf("User not found %s ", err.Error()), codes.NotFound)
	}
	// check if the device is already present
	for _, device := range cwUser.GetLinkedDevices() {
		if device.GetTvEmac() == request.GetLinkedDevice().GetTvEmac() {
			return nil, makingError("Database error", "Device already linked ", codes.AlreadyExists)
		}
	}
	cwUser.LinkedDevices = append(cwUser.LinkedDevices, request.LinkedDevice)
	_, err = h.ReplaceOne(ctx, bson.M{"googleid": request.GetGoogleId()}, cwUser)
	if err != nil {
		response.IsLinkedDeviceFetched = false
		return &response, makingError("Database error", fmt.Sprintf(err.Error()), codes.Internal)
	}
	response.LinkedDevices = cwUser.LinkedDevices
	response.IsLinkedDeviceFetched = true
	return &response, nil
}

func (h *Server) RemoveTvDevice(ctx context.Context, request *pb.RemoveTvDeviceRequest) (*pb.RemoveTvDeviceResponse, error) {
	var (
		cwUser   pb.User
		response pb.RemoveTvDeviceResponse
		err      error
	)
	if err = h.FindOne(context.Background(), bson.M{"googleid": request.GetGoogleId()}).Decode(&cwUser); err != nil {
		return nil, makingError("Database error", fmt.Sprintf("Error while fetching user data %s ", err.Error()), codes.Internal)
	}
	for i, v := range cwUser.LinkedDevices {
		if v.TvEmac == request.TvEmac {
			cwUser.LinkedDevices = append(cwUser.LinkedDevices[:i], cwUser.LinkedDevices[i+1:]...)
			_, err = h.ReplaceOne(ctx, bson.M{"googleid": request.GetGoogleId()}, cwUser)
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
	var (
		cwUser   pb.User
		err      error
		response pb.LinkedDeviceResponse
	)
	if err = h.FindOne(ctx, bson.D{{"googleid", request.GoogleId}}).Decode(&cwUser); err != nil {
		return nil, makingError("Database error", fmt.Sprintf("Error while fetching user data %s ", err.Error()), codes.Internal)
	}
	response.LinkedDevices = cwUser.LinkedDevices
	response.IsLinkedDeviceFetched = true
	return &response, nil
}

func (h *Server) GetLinkedUsers(ctx context.Context, linkUserReq *pb.LinkUserRequest) (*pb.LinkUserResponse, error) {
	var (
		err         error
		linkedUsers pb.LinkUserResponse
	)
	// find all the user who have the requested tv emac
	resultCursor, err := h.Find(ctx, bson.M{"linkeddevices.tvemac": linkUserReq.GetTvEmac()})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, makingError("Database Error", fmt.Sprintf("No Users Linked to %s TV emac", linkUserReq.GetTvEmac()), codes.NotFound)
		}
		return nil, makingError("Database Error", fmt.Sprint(err.Error()), codes.Internal)
	}
	defer resultCursor.Close(ctx)
	var user *pb.User
	for resultCursor.Next(ctx) {
		err = resultCursor.Decode(user)
		if err != nil {
			return nil, makingError("Internal Error", fmt.Sprintf("Error while decoding data %s", err.Error()), codes.Internal)
		}
		linkedUsers.User = append(linkedUsers.User, user)
	}
	return &linkedUsers, nil
}


func makingError(filedName, desc string, code codes.Code) error {
	return status.Error(code, desc)
}