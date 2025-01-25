
run:
	export ADDR=":9999" && go run ./cmd/web/ -addr=$$ADDR
test:
	go test -v ./cmd/web/
push:
	git push github
	git push codeberg 
