.PHONY: model build test clean

model:
	GO111MODULE=on go build -o model cmd/modelgen/modelgen.go

build:
	GO111MODULE=on go build -o main cmd/gormgen/gormgen.go

test:
	go test -v ./...

clean:
	@rm -rf model main