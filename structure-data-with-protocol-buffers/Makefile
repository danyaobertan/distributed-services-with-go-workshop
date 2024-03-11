install:
	go mod tidy
	
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=.

test:
	CGO_ENABLED=1 go test -race ./...


