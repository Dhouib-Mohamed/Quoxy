# check if go dependecies are installed
if [ ! -d "vendor" ]; then
    echo "Installing go dependencies"
    go mod vendor
fi

# start the server token-management
echo "Starting server token-management"
cd ..
go run cmd/token-proxy/main.go
