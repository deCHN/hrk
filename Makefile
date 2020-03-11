tag:
	go mod tidy
	go test ./...
	git status
	git tag v0.1.0
	git push origin v0.1.0
