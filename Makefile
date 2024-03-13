test:
	go test  -count=1 ./...

coverage:
	go test -count=1 ./... -cover