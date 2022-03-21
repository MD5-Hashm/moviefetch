@echo off

set GOOS=windows
set GOARCH=amd64
go build -o "win-amd64-moviefetch.exe"
echo "BUILT: windows-amd64"

set GOOS=linux
set GOARCH=amd64
go build -o "linux-amd64-moviefetch"
echo "BUILT: linux-amd64"

set GOOS=darwin
set GOARCH=amd64
go build -o "osx-amd64-moviefetch"
echo "BUILT: osx-amd64"

set GOOS=darwin
set GOARCH=arm64
go build -o "osx-arm64-moviefetch"
echo "BUILT: osx-arm64"

echo "DONE"