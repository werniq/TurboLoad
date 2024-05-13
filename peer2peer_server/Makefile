build:
	@echo "[*] BUILDING DIRECTORY"
	go build -o bin/go100mbps.exe .

run: build
	@echo "[*] RUNNING SERVER"
	./bin/go100mbps.exe

no-build:
	@echo "[*] RUNNING SERVER (WITHOUT BUILDING)"
	./bin/go100mbps

test:
	@echo "[*] TESTING GO FILES"
	go test ./... -v