# Structure Data with Protocol Buffers

> In this chapter, we covered the protobuf fundamentals we’ll use throughout
our project. These concepts will be vital throughout our project, especially as
we build our gRPC client and server. Now let’s create the next vital piece of
our project: a commit log library.

## Install Protocol Buffers

To install Protocol Buffers on Ubuntu, run the following command:

```shell
sudo apt install protobuf-compiler
```
To install the go protobuf runtime:

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Compile Protocol Buffers

```shell
make compile
```

