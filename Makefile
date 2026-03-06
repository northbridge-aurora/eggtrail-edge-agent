BINARY_NAME=eggtrail-agent

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test ./...

attest:
	@echo "Generating in-toto attestation for v2.6.1-spring..."
	# Logic handled by .github/workflows/handoff.yml
