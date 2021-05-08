FROM scratch
COPY webserver-go /
COPY html /html
ENTRYPOINT ["/webserver-go"]