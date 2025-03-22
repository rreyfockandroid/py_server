#protoc-gen:
#	./protoc/bin/protoc --proto_path=./ --python_out=./ api.proto

proto-install:
	todo

protoc-gen:
	protoc --go_out=. --go-grpc_out=. protos/api.proto

proto-gen:
	python -m grpc_tools.protoc -I./protos --python_out=. --pyi_out=. --grpc_python_out=. api.proto

server-start:
	python server_asyncio.py

server-stop:
	ps aux | grep 'server_asyncio.py' | grep -v 'grep' | tr -s ' ' '^' | cut -d '^' -f 2 | xargs kill -15

server-w-start:
	python server.py

client-run:
	python client.py