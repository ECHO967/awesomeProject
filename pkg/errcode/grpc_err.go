package errcode

import (
	pb "awesomeProject/pkg/proto/err"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToRPCError(err *Error) error {
	pbErr := &pb.Error{Code: int32(err.Code()), Message: err.Msg()}
	s, _ := status.New(codes.Unknown, err.Msg()).WithDetails(pbErr)
	return s.Err()
}

type Status struct {
	*status.Status
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(codes.Unknown, msg).WithDetails(&pb.Error{Code: int32(code), Message: msg})
	return &Status{s}
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}
