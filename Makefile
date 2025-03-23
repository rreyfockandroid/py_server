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
	ps aux | egrep 'server_asyncio.py|server.py' | grep -v 'grep' | tr -s ' ' '^' | cut -d '^' -f 2 | xargs kill -15

server-w-start:
	python server.py

client-run:
	python client.py

docker-build:
	docker build -t py-server-img .

docker-run:
	docker run -d --name py_server -p50051:50051 py-server-img

docker-stop:
	docker stop py_server

docker-start:
	docker start py_server

docker-logs:
	docker logs -f py_server

docker-rm:
	docker rm py_server