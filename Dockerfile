FROM scratch

COPY bin/hello hello

EXPOSE 8000

ENTRYPOINT ["/hello"]
