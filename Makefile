gen:
	buf generate -v --template api/grpc/buf.gen.yaml

sqlc:
	sqlc generate -f internal/storage/postgres/sqlc.yml


.PHONY: gen sqlc