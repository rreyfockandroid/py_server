FROM python:3.10.12-slim

WORKDIR /app

ADD server.py .
ADD utils/ utils/
ADD datasource/ datasource/
ADD api_pb2_grpc.py .
ADD api_pb2.py .
ADD api_pb2.pyi .
ADD requirements.txt .

RUN pip install -r requirements.txt

CMD ["python","-u","server.py"]
