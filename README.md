# Go gRPC Server Example

This code just a simple example for creating gRPC Server on GO, for test this server just use [go-grpc-client-example](https://github.com/rendyuwu/go-grpc-client-example)

This example just return simple user data
- username
- name
- age

And the list of slice can be found at user/delivery/grpc/user_handler.go

At this sample we just have 2 valid username
- rendyuwu
- budi

Because this just example for handling gRPC request, so we don't implement database, data simply on declared slice

### How To Run
- go run main.go (port used in this sample is ":1200")