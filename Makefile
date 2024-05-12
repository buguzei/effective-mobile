.PHONY:
repo-mocks:
	mockgen -source=auth/internal/repo/repo.go -destination=auth/internal/repo/mocks/repo_mock.go

auth-proto:
	protoc -I proto proto/auth/auth.proto --go_out=./pkg/protos/gen/auth/ --go_opt=paths=source_relative --go-grpc_out=./pkg/protos/gen/auth/ --go-grpc_opt=paths=source_relative

cars-proto:
	protoc -I ./pkg/protos/proto ./pkg/protos/proto/cars/cars.proto --go_out=./pkg/protos/gen/ --go_opt=paths=source_relative --go-grpc_out=./pkg/protos/gen/ --go-grpc_opt=paths=source_relative