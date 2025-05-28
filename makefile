run:
	go run main.go

run-script:
	go run main.go -- script.txt


build:
	go build -o inky main.go

test:
	go test -v

test-e2e:
	go test -v -run TestE2E