FROM scratch
COPY webserver-go /
COPY static /static
ENTRYPOINT ["/webserver-go"]