run:
	go run ./cmd/api/

dev:
	nodemon --exec go run ./cmd/api/ --signal TERM