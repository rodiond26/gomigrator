# postgres:
# 	docker run --rm -ti --network host -e POSTGRES_PASSWORD=postgres postgres
postgres:
	docker run --rm -ti --name database --network host -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres

# adminer:
# 	docker run --rm -ti --network host adminer

# migrate:
# 	migrate -source file://migrations \
# 			-database postgres://postgres:secret@localhost/postgres&sslmode=disable up

.PHONY: postgres # adminer migrate
