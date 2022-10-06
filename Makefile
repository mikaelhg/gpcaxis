.PHONY: clean test build

build:
	go build -ldflags="-s -w" -o ./bin/pcaxis2parquet ./cmd/pcaxis2parquet/

cross:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-linux-amd64 ./cmd/pcaxis2parquet/
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-linux-arm64 ./cmd/pcaxis2parquet/
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-darwin-amd64 ./cmd/pcaxis2parquet/
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-darwin-arm64 ./cmd/pcaxis2parquet/
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-windows-amd64.exe ./cmd/pcaxis2parquet/
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" \
		-o ./bin/pcaxis2parquet-windows-arm64.exe ./cmd/pcaxis2parquet/

clean:
	go clean
	rm -f ./bin/pcaxis2parquet
	rm -f ./bin/pcaxis2parquet-linux-amd64
	rm -f ./bin/pcaxis2parquet-linux-arm64
	rm -f ./bin/pcaxis2parquet-darwin-amd64
	rm -f ./bin/pcaxis2parquet-darwin-arm64
	rm -f ./bin/pcaxis2parquet-windows-amd64.exe
	rm -f ./bin/pcaxis2parquet-windows-arm64.exe

test:
	time -v go run ./cmd/pcaxis2parquet/main.go \
		--file=./data/statfin_altp_pxt_12bd.px

test2:
	time -v ./bin/pcaxis2parquet --file=./data/statfin_vtp_pxt_124l.px
