# Dependencies:
# Yarn
# Go Toolchain 1.13 or newer
#   go: github.com/markbates/pkger/cmd/pkger
#   go: src.techknowlogick.com/xgo
# docker
#   docker: techknowlogick/xgo:latest
# upx

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
		pkger &&\
		xgo -tags $(RELEASETAGS) -ldflags $(LDFLAGS) -out $(OUTPUTPREFIX) -targets "linux/amd64,darwin-10.11/amd64,windows-7.0/386" . &&\
		upx $(UPXFLAGS) $(OUTPUTPREFIX)-* 
	mv backend/$(OUTPUTPREFIX)-* .
	chmod +x $(OUTPUTPREFIX)-*
