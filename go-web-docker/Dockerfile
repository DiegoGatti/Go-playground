FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .    
RUN go build -o /out/example .
FROM scratch as bin
COPY --from=build /out/example /