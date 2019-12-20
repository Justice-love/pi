//go:generate protoc -I=../../api --go_out=paths=source_relative,plugins=grpc:. car.proto

package car
