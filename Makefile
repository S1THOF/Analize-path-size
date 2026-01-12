build: 
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run: build
	bin/hexlet-path-size testdata/file1.csv

	
test:
	go test -v