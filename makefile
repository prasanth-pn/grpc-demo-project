export:
	export PATH="$PATH:$HOME/.local/bin"

run:
	go run main.go
protoc:
	protoc proto/*.proto --go_out=plugins=grpc:.
	#protoc --go_out=. --go_grpc_out=.  proto/greet.proto