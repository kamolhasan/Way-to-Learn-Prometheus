FROM busybox:glibc

COPY Way-to-Learn-Prometheus /bin/api

ENTRYPOINT ["/bin/api"]

EXPOSE 8080