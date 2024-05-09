export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://admin:admin@localhost/news?sslmode=disable

export-env:
	@export GOOSE_DRIVER=$(GOOSE_DRIVER);
	@export GOOSE_DBSTRING=$(GOOSE_DBSTRING);

goose-init: export-env
	@goose -s -dir=./migrations init

goose-create: export-env
	@read -p "Enter migration name: " migration_name; \
	goose -s -dir=./migrations create $$migration_name sql

goose-up: export-env
	@goose -dir=./migrations up

goose-down: export-env
	@goose -dir=./migrations down
