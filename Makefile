\.PHONY: help
SHELL := /bin/zsh

help: ## display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## build the project
	@echo "Building the project"
	go mod tidy
	go build -o conv main.go
	chmod +x conv

convert: ## converts original markdown file to mark compatible with html metadata
	@echo "Copying attachment files"
	cp -r input/images output/images
	@echo "Converting markdown files to html"
	@for file in input/*.md; do \
		filename=$$(basename $$file); \
		echo "Converting $$filename"; \
		./conv --input ./input/$$filename --output ./output/$$filename; \
	done

push: ## pushing converted files to the confluence using mark
	@echo "Pushing files to confluence"
	@for file in output/*.md; do \
		filename=$$(basename $$file); \
		echo "Pushing $$filename"; \
		docker run --rm -i -v .:/workspace -w /workspace kovetskiy/mark:latest \
			mark -c config.ini --files ./output/$$filename; \
	done
