# Write a Log Package

> You now know what logs are, why they’re important, and how they’re used in various applications including distributed services. And then you learned how  to  build  one!  This  log  serves  as  the  foundation  of  our  distributed  log. Now we can build a service on our library and make the library’s functionality accessible to people on other computers.

## Install Protocol Buffers

To install Protocol Buffers on Ubuntu, run the following command:

```shell
sudo apt install protobuf-compiler
```

To install the go grpc runtime:
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

To install the go protobuf runtime:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Compile Protocol Buffers

```shell
make compile
```