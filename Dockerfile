FROM golang:1.20

RUN mkdir /app
COPY . /app
WORKDIR /app

ENV PORT $PORT
EXPOSE $PORT

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl -f http://localhost/health || exit 1

RUN go build -o go-server .
CMD ["/app/go-server"]