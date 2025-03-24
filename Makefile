upmod:
	go get -u ./...
	go mod tidy
tag:
	sh scripts/tag.sh $(shell cat ./version)