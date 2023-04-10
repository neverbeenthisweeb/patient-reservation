FROM golang:1.18
WORKDIR /goapp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./patientreservation ./

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /goapp
COPY --from=0 /goapp/patientreservation ./
CMD ["./patientreservation"]