NAME     := mining

.PHONY: build
build:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif
	glide install
	go build -o cmd/$(NAME) mining.go

.PHONY: clean
clean:
	rm -rf cmd/*
