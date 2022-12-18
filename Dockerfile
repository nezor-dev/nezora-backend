FROM golang:1.20rc1-alpine3.16 AS builder

WORKDIR /build

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch

# Copy binary and config files from /build 
# to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Export necessary port.
EXPOSE 3001

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]