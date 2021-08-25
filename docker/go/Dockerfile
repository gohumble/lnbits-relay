FROM golang:1.17-alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN echo -e  "\e[91mSTART BUILDING PROCESS ON \c" && go version && echo -e "\e[0m"

RUN go build -o relay .

#building finished. Now extracting single bin in second stage.
FROM alpine

EXPOSE 8081
ENV HTTP_PROXY=tor:8110
COPY --from=builder /build/relay /app/
COPY --from=builder /build/config.yaml /app/

WORKDIR /app

CMD ["./relay"]