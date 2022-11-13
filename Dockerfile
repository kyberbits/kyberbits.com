FROM node:18-buster as parcelBuilder
WORKDIR /staging
COPY . .
RUN make build-npm

FROM aaronellington/valet:latest
COPY --from=parcelBuilder /staging/var/dist .
EXPOSE 1234
