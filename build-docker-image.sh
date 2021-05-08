CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

BIN=./webserver-go
if [ -e $BIN -a -f $BIN ]; then
    echo "Build image for $BIN"
    docker build -t ityuhui/webserver-go . \
	--network host \
	--build-arg HTTP_PROXY=$http_proxy \
	--build-arg HTTPS_PROXY=$http_proxy \
	--build-arg http_proxy=$http_proxy \
	--build-arg https_proxy=$http_proxy
else
    echo "$BIN does not exist."
fi
