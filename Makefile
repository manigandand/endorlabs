lint:
	./golangci.sh ./... $(LINT_ENV)
	# go mod tidy -compat=1.17

gotests:
	go test -v --failfast ./...

test: gotests

build: build_gobin
	# docker build -t endorlabs:latest .

build_gobin:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o endorlabs -ldflags="-X 'main.VersionLabel=$(versionLabel)'" .

run: build
	./endorlabs config.json
