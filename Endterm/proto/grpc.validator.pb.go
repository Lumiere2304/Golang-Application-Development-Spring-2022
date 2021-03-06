package example

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	time "time"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

func (this *User) Validate() error {
	if !(this.ID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`ID must a positive integer`))
	}
	if this.CreateDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreateDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreateDate", err)
		}
	}
	return nil
}
func (this *UserRole) Validate() error {
	return nil
}
func (this *UpdateUserRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.UpdateMask != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateMask); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateMask", err)
		}
	}
	return nil
}
func (this *ListUsersRequest) Validate() error {
	if this.CreatedSince != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedSince); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedSince", err)
		}
	}
	if this.OlderThan != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OlderThan); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OlderThan", err)
		}
	}
	return nil
}