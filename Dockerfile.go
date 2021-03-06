FROM golang as base
WORKDIR /app
COPY main.go /app
RUN go build main.go

FROM alpine
WORKDIR /app
COPY --from=base /app/main /app 
CMD ["./main"]


