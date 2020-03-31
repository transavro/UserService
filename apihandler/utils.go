package apihandler

import (
	pb "UserService/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"strconv"
	"strings"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var IsNumber = regexp.MustCompile(`^[0-9]+$`).MatchString

func validatingData(user *pb.User) error {
	// checking UserName
	//checking if the userName is too small
	if len(user.GetName()) <= 3 {
		// if the username is too small
		return makingError("Username", "Username  must be greater than 3", codes.InvalidArgument)
	} else if !IsLetter(strings.ReplaceAll(user.GetName(), " ", "") ){
		return makingError("Username", "Username  must contain alphanumeric characters", codes.InvalidArgument)
	}

	// checking Genre
	if user.GetGenre() == nil || len(user.GetGenre()) == 0 {
		return makingError("User Genres", "User Genres could not be empty", codes.InvalidArgument)
	} else if user.GetLanguage() == nil || len(user.GetLanguage()) == 0 {
		return makingError("User Language", "User Language could not be empty", codes.InvalidArgument)
	} else if user.GetContentType() == nil || len(user.GetContentType()) == 0 {
		return makingError("User Categories", "User Categories could not be empty", codes.InvalidArgument)
	}

	// checking phone number
	if len(user.GetPhoneNumber()) == 0 {
		return makingError("Phone Number", "Could not be empty", codes.InvalidArgument)
	} else if !IsNumber(user.GetPhoneNumber()) {
		return makingError("Phone Number", "Must contain only numbers", codes.InvalidArgument)
	} else if len(user.GetPhoneNumber()) != 10 {
		return makingError("Phone Number", "Must contain 10 numbers", codes.InvalidArgument)
	}

	// check if the Data is valid
	if len(user.GetDateOfBirth()) == 0 {
		return makingError("Date of Birth", "Cannot be empty", codes.InvalidArgument)
	} else {
		// validating DOB
		dateSlipt := strings.Split(user.GetDateOfBirth(), "/")
		if len(dateSlipt) != 3 && len(dateSlipt[0]) != 2 && len(dateSlipt[1]) != 2 && len(dateSlipt[2]) != 4 {
			return makingError("Date of Birth", "Format is invalid. valid format is DD/MM/YYYY", codes.InvalidArgument)
		} else {
			u, err := strconv.ParseUint(dateSlipt[2], 10, 64)
			if err != nil {
				return makingError("Date of Birth", "Error while decoding Date of Birth", codes.Internal)
			}
			if u > 2007 {
				return makingError("Date of Birth", "User is too young in register", codes.InvalidArgument)
			}
		}
	}
	return nil
}

func makingError(filedName, desc string, code codes.Code) error {
	return status.Error(code, desc)

	//st := status.New(code, fmt.Sprintf("invalid %s", filedName))
	//v := &errdetails.BadRequest_FieldViolation{
	//	Field:       filedName,
	//	Description: desc,
	//}
	//br := &errdetails.BadRequest{}
	//br.FieldViolations = append(br.FieldViolations, v)
	//st, err := st.WithDetails(br)
	//if err != nil {
	//	panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	//}
	//return st.Err()
}
