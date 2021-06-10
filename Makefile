#
#
#

proto:
	GOBIN=/bin PATH=$$GOBIN:$$PATH protoc --go_out=. --twirp_out=. api/billing.proto

build:
	cd cmd/billing;	go build .; mv billing ../../bin; cd ../..
