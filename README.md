# grpc-example



### References



### Protobuf source generator
``` sh
cd proto-files

protoc --go_out=../gen --go_opt=paths=source_relative --go-grpc_out=../gen  --go-grpc_opt=paths=source_relative ./vehicle.proto  
```

