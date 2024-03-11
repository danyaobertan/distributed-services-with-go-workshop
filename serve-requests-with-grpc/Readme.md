# Serve Requests with gRPC

> You now know how to define a gRPC service in protobuf, compile your gRPC protobufs into code, build a gRPC server, and test that everything works end-to-end across your client and server. You can build a gRPC server and client, and you can use your log over the network. Next  we’ll  improve  the  security  of  our  service  by  encrypting  the  data  sent between the client and server with SSL/TLS, and authenticating requests so
we can know who’s making each request and whether they’re allowed to.

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

Add the following to your .bashrc or .zshrc file:
```shell
	export PATH=$PATH:$HOME/go/bin
```

## Compile Protocol Buffers


```shell
make compile
```