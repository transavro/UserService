// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: User.proto

package User

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/type/date"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *LinkUserRequest) Validate() error {
	if this.TvEmac == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvEmac", fmt.Errorf(`Tv Emac cannot be empty.`))
	}
	return nil
}
func (this *LinkUserResponse) Validate() error {
	for _, item := range this.User {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("User", err)
			}
		}
	}
	return nil
}
func (this *RemoveTvDeviceRequest) Validate() error {
	if this.GoogleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GoogleId", fmt.Errorf(`User Id cannot be empty.`))
	}
	if this.TvEmac == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvEmac", fmt.Errorf(`Tv Emac cannot be empty.`))
	}
	return nil
}
func (this *RemoveTvDeviceResponse) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	if this.GoogleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GoogleId", fmt.Errorf(`User Id cannot be empty.`))
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	if this.GoogleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GoogleId", fmt.Errorf(`User Id cannot be empty.`))
	}
	return nil
}
func (this *DeleteReponse) Validate() error {
	return nil
}
func (this *LinkedDeviceResponse) Validate() error {
	for _, item := range this.LinkedDevices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LinkedDevices", err)
			}
		}
	}
	return nil
}
func (this *TvDevice) Validate() error {
	if this.GoogleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GoogleId", fmt.Errorf(`User Id cannot be empty.`))
	}
	if this.LinkedDevice != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LinkedDevice); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LinkedDevice", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	return nil
}
func (this *CreateReponse) Validate() error {
	return nil
}

var _regex_User_Name = regexp.MustCompile(`^[a-zA-Z ]{3,20}$`)
var _regex_User_PhoneNumber = regexp.MustCompile(`^[0-9]*$`)

func (this *User) Validate() error {
	if !_regex_User_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`User name cannot contain Numbers`))
	}
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`Email id cannot contain be empty`))
	}
	if !_regex_User_PhoneNumber.MatchString(this.PhoneNumber) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`Phone number must be of 10 digit and should be consist of only numbers.`))
	}
	if this.PhoneNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`Phone number must be of 10 digit and should be consist of only numbers.`))
	}
	if !(len(this.PhoneNumber) == 10) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`Phone number must be of 10 digit and should be consist of only numbers.`))
	}
	if this.GoogleId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GoogleId", fmt.Errorf(`User Id cannot be empty.`))
	}
	if nil == this.DateOfBirth {
		return github_com_mwitkow_go_proto_validators.FieldError("DateOfBirth", fmt.Errorf("message must exist"))
	}
	if this.DateOfBirth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DateOfBirth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DateOfBirth", err)
		}
	}
	if _, ok := Gender_name[int32(this.Gender)]; !ok {
		return github_com_mwitkow_go_proto_validators.FieldError("Gender", fmt.Errorf(`Gender cannot be empty or Invalid.`))
	}
	if this.ImageUrl == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ImageUrl", fmt.Errorf(`Image Url cannot be empty.`))
	}
	if len(this.Genre) < 3 {
		return github_com_mwitkow_go_proto_validators.FieldError("Genre", fmt.Errorf(`Please Select Genre greater than 3 and cannot be empty.`))
	}
	for _, item := range this.Genre {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Genre", fmt.Errorf(`Please Select Genre greater than 3 and cannot be empty.`))
		}
	}
	if len(this.Language) < 3 {
		return github_com_mwitkow_go_proto_validators.FieldError("Language", fmt.Errorf(`Please Select Language greater than 3 and cannot be empty.`))
	}
	for _, item := range this.Language {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Language", fmt.Errorf(`Please Select Language greater than 3 and cannot be empty.`))
		}
	}
	if len(this.ContentType) < 3 {
		return github_com_mwitkow_go_proto_validators.FieldError("ContentType", fmt.Errorf(`Please Select ContentType greater than 3 and cannot be empty.`))
	}
	for _, item := range this.ContentType {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("ContentType", fmt.Errorf(`Please Select ContentType greater than 3 and cannot be empty.`))
		}
	}
	for _, item := range this.LinkedDevices {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LinkedDevices", err)
			}
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *LinkedDevice) Validate() error {
	if this.TvEmac == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvEmac", fmt.Errorf(`Tv Emac cannot be empty.`))
	}
	if this.TvPanel == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvPanel", fmt.Errorf(`Tv Panel cannot be empty.`))
	}
	if this.TvBoard == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvBoard", fmt.Errorf(`Tv Board cannot be empty.`))
	}
	if this.TvName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TvName", fmt.Errorf(`Tv Name cannot be empty.`))
	}
	return nil
}
