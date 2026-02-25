
echo "buding..."

apt install golang
go build -tag netgo -o flager main.go
echo "add to access"
chmod +x flager
echo "Now Yo can you flager tool"