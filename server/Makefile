.PHONY:cephweb
.PHONY:cephweb-win

cephweb:
	go build -tags luminous -o cephweb
cephweb-win:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -tags luminous -ldflags "-linkmode external -extldflags -static -s -w" -o cephweb.exe
