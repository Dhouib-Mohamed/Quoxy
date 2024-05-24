# check if go dependecies are installed
if [ ! -d "vendor" ]; then
    echo "Installing go dependencies"
    go mod vendor
fi

cd ..

# Create named pipes
mkfifo pipe-token pipe-proxy

# Command 1
cmd1="go run cmd/token-management/main.go"
# Command 2
cmd2="go run cmd/proxy/main.go"

# Run commands in parallel, redirecting their outputs to the named pipes
{
  $cmd1 > pipe-token
} &

{
  $cmd2 > pipe-proxy
} &

# Read from the pipes in parallel and prefix output with command identifiers
{
  while read -r line; do
    echo "Token Management: $line"
  done < pipe-token
} &

{
  while read -r line; do
    echo "Proxy Server: $line"
  done < pipe-proxy
} &

# Wait for both commands to complete
wait

# Clean up named pipes
rm pipe-token pipe-proxy

