FROM golang:1.22-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download


RUN GOARCH=arm64 go build -o /service-cicd


FROM alpine
COPY --from=builder /service-cicd .


EXPOSE 80
CMD [ "/service-cicd" ]