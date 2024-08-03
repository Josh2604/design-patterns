.PHONY: run runc

run:
	go run ./cmd/web/

runc:
	go run ./cmd/web -cache=true