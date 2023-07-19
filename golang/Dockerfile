FROM docker.io/golang:alpine as build

WORKDIR /opt/app-root
ENV GOPATH=/opt/app-root
COPY ./ src
WORKDIR /opt/app-root/src
RUN go build -o prom_exmpl

FROM scratch
WORKDIR /opt/app-root
COPY --from=build /opt/app-root/src/prom_exmpl /opt/app-root/prom_exmpl
EXPOSE 2112
ENTRYPOINT ["./prom_exmpl"]
