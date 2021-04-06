FROM golang:latest as builder
WORKDIR /go/src/web
RUN go get github.com/gin-gonic/gin
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix temp -ldflags '-extldflags "-static"' .


FROM busybox
WORKDIR /root/
COPY --from=builder /go/src/web/web .
ADD html/index.html html/
ADD stylesheet stylesheet/
ADD resources /resources
ADD images images/
CMD ["./web"]

