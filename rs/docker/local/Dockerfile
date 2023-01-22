FROM golang:1.19-buster

WORKDIR /go/src/rs

COPY go.* ./
RUN go mod download
COPY . ./

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && go install github.com/volatiletech/sqlboiler/v4 \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql \
  && go install github.com/golangci/golangci-lint/cmd/golangci-lint \
  && go install github.com/99designs/gqlgen \
  && go install github.com/golang/mock/mockgen \
