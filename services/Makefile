all: get_vendor_deps install

get_vendor_deps:
	go get github.com/Masterminds/glide
	glide install
	
test:
	@go test `glide novendor`
	
install:
	go install ./cmd/explorercli
