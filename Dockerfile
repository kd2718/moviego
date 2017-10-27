FROM golang:alpine 
#RUN mkdir /app 
RUN apk update
RUN apk add curl
RUN apk add libcurl
RUN apk add libssh2
ADD . /go/src/github.com/kd2718/goplay
#VOLUME [ "/app" ]
#RUN curl https://glide.sh/get | sh
WORKDIR /go/src/github.com/kd2718/goplay 
RUN go build -o main . 
CMD ["/go/src/github.com/kd2718/goplay/main"]

#RUN glide update

# remote volums:  docker run -i -t -v /$(pwd)://go/src/github.com/kd2718/goplay goplay:glide