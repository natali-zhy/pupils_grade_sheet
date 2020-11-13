migrate-up:
	migrate -path db/migration -database "postgresql://postgres:@localhost:5432/students?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://postgres:@localhost:5432/students?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdblf