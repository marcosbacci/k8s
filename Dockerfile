FROM scratch

COPY /go/bin/hello /go/bin/hello

EXPOSE 8000

ENTRYPOINT ["/go/bin/hello"]
