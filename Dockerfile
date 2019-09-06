# builder image, uses all go buildchain deps
FROM golang:alpine as builder

# install deps
RUN apk --no-cache add build-base git bzr gcc

# set working directory and copy all files over
WORKDIR /build
COPY . .

# add go modules on so that we can create outside of $GOPATH
ENV GO111MODULE=on
RUN go build -o giaa .

# runner image, contains only runtime reqs
FROM alpine as runner

WORKDIR /app
COPY --from=builder /build/giaa .
ENTRYPOINT [ "./giaa" ]
