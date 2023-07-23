FROM docker.io/library/golang:1.20.2 AS build

#Unset the GOPATH
ENV GOPATH=

#Change the directory
WORKDIR /go

#Copy all files apart from the ones in .dockerignore
COPY . .

#Check if any modules need downloading
RUN go mod download

#Compile the golang code, CGO_ENABLE=0 removes cross compile dependencies
RUN CGO_ENABLED=0 go build

FROM scratch
#Add repository label
LABEL org.opencontainers.image.source "https://github.com/cloudziu/debugging-scratch"
#Add image name as a label
LABEL org.opencontainers.image.base.name "scratch"

#Copy the binaries to empty base image
COPY --from=build /go/faultyapp /bin/faultyapp

#Run server
WORKDIR /bin
ENTRYPOINT [ "./faultyapp" ]