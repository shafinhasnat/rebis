FROM golang:1.24-rc-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go mod download
RUN go build -o /rebis

FROM scratch
COPY --from=build /rebis /rebis
EXPOSE 6666
CMD ["/rebis"]
