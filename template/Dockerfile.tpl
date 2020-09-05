FROM golang:1.15
ARG version
WORKDIR work
COPY go.* .
COPY ../pkg pkg/
COPY ../vendor vendor/
COPY services/__SVC_NAME__/ services/__SVC_NAME__/
RUN GO111MODULE=on GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X __PACKAGE__/services/__SVC_NAME__/main.version=$version" -a -installsuffix cgo -o __SVC_NAME__ .

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /__SVC_NAME__/
COPY --from=0 work/services/__SVC_NAME__ .
ENTRYPOINT ["/__SVC_NAME__/__SVC_NAME__"]
CMD ["run", "--logtostderr"]
