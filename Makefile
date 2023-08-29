.PHONY : proto
proto:
	protoc proto/*.proto --go_out=proto/ --go-grpc_out=proto/

.PHONY : run
run:
	docker compose up --build

.PHONY : stop
stop:
	docker compose down