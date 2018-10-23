rm build_win64/*.exe
cp -rf assets/ build_win64/
env CGO_ENABLED="1" CC="/usr/bin/x86_64-w64-mingw32-gcc" GOOS="windows" CGO_LDFLAGS="-lmingw32 -lSDL2" CGO_CFLAGS="-D_REENTRANT" go build -o build_win64/DoomSlayeRL.exe *.go
 
