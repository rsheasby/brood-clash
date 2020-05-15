# Dependencies:
# Yarn
# Go Toolchain 1.13 or newer
#   go: github.com/markbates/pkger
#   go: github.com/techknowlogick/xgo
# docker
#   docker: techknowlogick/xgo:latest
# upx

LDFLAGS = "-s -w"
OUTPUTPREFIX = "brood-clash"

build:
	cd client &&\
		yarn &&\
		yarn build
	rm -f brood-clash-*
	rm -rf backend/static
	cp -r client/public backend/static
	cd backend &&\
		pkger &&\
		xgo -ldflags $(LDFLAGS) -out $(OUTPUTPREFIX) -targets "linux/amd64,darwin-10.11/amd64,windows-7.0/386" . &&\
		upx $(OUTPUTPREFIX)-* 
	mv backend/$(OUTPUTPREFIX)-* .

dev-backend:
	BCDEV=1 cd backend && go run .
