# Binary name
BINARY=hhsh
VERSION=v0.0.3

# Builds the project
build:
		GO111MODULE=on go build -o ${BINARY}
		GO111MODULE=on go test -v
# Installs our project: copies binaries
install:
		GO111MODULE=on go install
release:
		# Clean
		go clean
		rm -rf *.gz
		# Build for linux
		go clean
		GO111MODULE=on go build -o ${BINARY}
		tar czvf ${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}
		# Build for mac
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}
		tar czvf ${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}
		# Build for win
		go clean
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe
		tar czvf ${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe
		go clean
# Cleans our projects: deletes binaries
clean:
		go clean
		rm -rf *.gz
		rm hhsh*

.PHONY:  clean build