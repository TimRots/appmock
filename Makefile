appmock:
	GOOS=$(go env GOOS)
	GOARCH=$(go env GOARCH)
	BUILDNAME="appmock"
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ~/bin/${BUILDNAME} ./...

.PHONY: appmock
