# Dependencies:
# Yarn
# Go Toolchain 1.16 or newer

LDFLAGS = "-s -w"
RELEASETAGS = "release"
OUTPUTPREFIX = "brood-clash"
UPXFLAGS = --lzma --best

build:
	cd client &&\
		yarn &&\
		yarn build
	rm -f brood-clash-*
	rm -rf backend/static
	cp -r client/public backend/static
	cd backend &&\
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags $(RELEASETAGS) -ldflags $(LDFLAGS) -o $(OUTPUTPREFIX)-linux-amd64 &&\
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags $(RELEASETAGS) -ldflags $(LDFLAGS) -o $(OUTPUTPREFIX)-macos-amd64 &&\
		CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -tags $(RELEASETAGS) -ldflags $(LDFLAGS) -o $(OUTPUTPREFIX)-windows-i386.exe
	mv backend/$(OUTPUTPREFIX)-* .
	chmod +x $(OUTPUTPREFIX)-*
	upx $(UPXFLAGS) $(OUTPUTPREFIX)-*
