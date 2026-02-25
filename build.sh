GOOS=linux GOARCH=amd64 go build -tags netgo -o ./dist/flager main.go
tar -cJf ./dist/lnux-amd64-flager.tar.xz  ./dist/flager
rm -rf ./dist/flager
echo "build amd64 ✔️"
GOOS=linux GOARCH=arm go build -tags netgo -o ./dist/flager main.go
tar -cJf ./dist/lnux-arm-flager.tar.xz  ./dist/flager
rm -rf ./dist/flager
echo "build arm ✔️"
GOOS=linux GOARCH=386 go build -tags netgo -o ./dist/flager main.go
tar -cJf ./dist/lnux-386-flager.tar.xz  ./dist/flager
rm -rf ./dist/flager
echo "build 386 ✔️"
