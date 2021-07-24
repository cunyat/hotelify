FROM golang:1.16

RUN go get github.com/cespare/reflex
COPY reflex.conf /
COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
