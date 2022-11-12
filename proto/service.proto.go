syntax = "proto3"

package proto

option go_package = "proto;service"

message Request {
int64 a = 1;
int64 b = 2;
}

message Response {
int64 result = 1;
}

service AddService {
rpc Add(Request) returns (Response);
rpc Multiply(Request) returns (Response);
}