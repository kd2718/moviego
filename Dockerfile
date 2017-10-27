FROM golang:latest 
RUN mkdir /app 
#ADD . /app/
#VOLUME [ "/app" ]
#WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]

RUN curl https://glide.sh/get | sh
#RUN glide update