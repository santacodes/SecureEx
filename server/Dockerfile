FROM golang:1.20

WORKDIR /usr/app

RUN export GO111MODULE=on
RUN cd /usr/app && git clone https://github.com/santacodes/SecureEx.git

<<<<<<< Updated upstream
RUN cd /usr/app/SecureEx/server && go build
=======
RUN go mod download

COPY *.go ./
COPY *.db ./
COPY *.env ./

RUN go build -o /docker-gs-ping
>>>>>>> Stashed changes

EXPOSE 3000

CMD [ "/usr/app/SecureEx/server/server" ]
