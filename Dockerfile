FROM scratch
ADD webserver-go /
ENTRYPOINT ["/webserver-go"]