.PHONY: run build test clean deps lint fmt vet

# Run the service directly (fast iteration. no binary artifact)
run:
		go run ./cmd/service/*

# Tidy and download deps
deps:
		go mod tidy
		go mod download

# Clean build artifacts
clean:
		rm -rf bin/ coverage.out
