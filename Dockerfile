FROM golang:1.24.6-alpine AS build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /gshopping

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /gshopping /gshopping
COPY --from=build /app/db/migrations /db/migrations
COPY --from=build /app/config.yml /config.yml
COPY --from=build /app/banner.txt /banner.txt
EXPOSE 4000
USER nonroot:nonroot

ENTRYPOINT ["/gshopping"]
