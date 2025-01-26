
run:
	export ADDR=":9999" && go run ./cmd/web/ -addr=$$ADDR
test:
	go test -v ./cmd/web/
test-all:
	go test -race ./...
test-perf:
	go test -count=100 ./...
push:
	git push github
	git push codeberg 
config:
	go mod tidy
	mkdir tls
	go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost --outdir=./tls