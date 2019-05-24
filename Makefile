LISTGOFILES=$$(go list ./... | grep -v cmd)

test: 
	DATA_HOME=$$PWD/testdata \
	go test -count=1 -v $(LISTGOFILES)

run:
	# export CONF_HOME=$$PWD/conf
	DATA_HOME=$$PWD/testdata \
	go run cmd/matrix-portal/main.go

check:
	test -z $$(go fmt $(LISTGOFILES)) 
	go vet ./...
	golint -set_exit_status $(LISTGOFILES)
	staticcheck ./...
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o gitsync ./cmd/gitsync-cli