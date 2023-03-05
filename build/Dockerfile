FROM golang:1.20 AS build
WORKDIR /app
COPY . ./
RUN make build

FROM alpine
COPY --from=build /app/bin/image-previewer /image-previewer
COPY --from=build /app/.env /.env

RUN chmod a+x image-previewer
CMD ["./image-previewer"]