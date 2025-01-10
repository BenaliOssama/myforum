run:
	export ADDR=":9999" && go run cmd/web/* -addr=$$ADDR

push:
	git push origin