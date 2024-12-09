test: 
	mkdir db
	go run cmd/main.go
push:
	git push github 
	git push codeberg
