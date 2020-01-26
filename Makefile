install:
	go build -o kcs cmd/main.go
	mv kcs /usr/local/bin/