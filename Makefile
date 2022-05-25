.PHONY: model build clean

model:
	GO111MODULE=on go build -o model cmd/modelgen/modelgen.go

build:
	GO111MODULE=on go build -o main cmd/gormgen/gormgen.go

clean:
	@rm -rf model main