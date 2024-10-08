# See https://github.com/chemidy/smallest-secured-golang-docker-image/blob/master/docker/scratch.Dockerfile

############################
# STEP 1 build executable binary
############################
FROM golang:1.23-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create user
ENV USER=sniper
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"
WORKDIR $GOPATH/src/github.com/RagOfJoes/findthesniper.io/

# Copy go modules then download and verify
COPY go.mod .
RUN go mod download
RUN go mod verify

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-w -s -extldflags "-static"' -a \
  -o /go/bin/findthesniper cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/findthesniper /go/bin/findthesniper

# Use an unprivileged user.
USER sniper:sniper

# Run the binary.
ENTRYPOINT ["/go/bin/findthesniper", "start"]
