FROM scratch
COPY webserver-go /
COPY public /public
ENTRYPOINT ["/webserver-go"]