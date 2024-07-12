ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: all build build-templ build-tailwind sqlc-generate migrate-diff watch

# Build the application
all: build

#build: build-templ build-tailwind
build: build-templ build-tailwind
	@echo "Building go binaries..."
	@go build -o bin/api/main cmd/api/main.go

build-templ:
	@echo "Building templ views..."
	@templ generate

build-tailwind:
	@echo "Building tailwindcss..."
	@tailwindcss -i static/assets/css/input.css -o static/assets/css/output.css

sqlc-generate:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate -f db/sqlc.yml

migrate-diff:
ifndef NAME
	$(error NAME is not set)
endif
	@atlas migrate diff -c file://db/atlas.hcl --env local $(NAME)

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi