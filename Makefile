GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
NAME = appmock

appmock:
	GO111MODULE=on GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${NAME} ./...

.PHONY: appmock
