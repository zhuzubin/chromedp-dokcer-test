#FROM golang:latest
#
#RUN mkdir /go/src/app
#
#ADD ./main.go /go/src/app
#COPY ./go.mod /go/src/app
#
#WORKDIR /go/src/app
#RUN go build
#
#CMD ["./app"]

# =============== build stage ===============
FROM golang:1.19.1 AS build

RUN mkdir -p /app

ADD ./main.go /app
COPY ./go.mod /app
COPY ./go.sum /app

WORKDIR /app

COPY go.* ./
RUN go mod download -x all

COPY . ./
# ldflags:
#  -s: disable symbol table
#  -w: disable DWARF generation
# run `go tool link -help` to get the full list of ldflags
RUN go build -ldflags "-s -w" -o chromedp-test -v main.go

# =============== final stage ===============
FROM chromedp/headless-shell:latest

WORKDIR /app
COPY --from=build /app/chromedp-test ./
ENTRYPOINT ["/app/chromedp-test"]