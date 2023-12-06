FROM golang:1.21

WORKDIR /opt/escrowsimple

COPY go.mod /opt/escrowsimple/
COPY go.sum /opt/escrowsimple/
RUN go mod download

COPY cmd/ /opt/escrowsimple/cmd/
COPY pkg/ /opt/escrowsimple/pkg/
RUN go build -o /opt/escrowsimple/bin/escrowsimple cmd/server/main.go

CMD [ "/opt/escrowsimple/bin/escrowsimple" ]
