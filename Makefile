all: site go
run: site go-run

go:
	go build cmd/bplog.go
go-run:
	go run cmd/bplog.go
site:
	$(MAKE) -C ui all