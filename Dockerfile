FROM golang:1.20-bullseye

ENV APPDIR=/workspace

RUN mkdir -p /workspace
WORKDIR $APPDIR

COPY go.mod go.sum $APPDIR/
RUN go mod download
COPY application $APPDIR/

RUN go install $APPDIR/

EXPOSE 8080

CMD [ "distributed-social-media" ]
