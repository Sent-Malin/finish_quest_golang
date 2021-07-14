include .env
export
migrateup: 
	migrate -source file\://go-migrate -database postgres\://$(user):$(password)@$(host)\:$(port)/$(dbname)?sslmode=disable up 1
migratedown: 
	migrate -source file\://go-migrate -database postgres\://$(user):$(password)@$(host)\:$(port)/$(dbname)?sslmode=disable down 1