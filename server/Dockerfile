FROM golang:1.20

WORKDIR /usr/app

RUN export GO111MODULE=on
RUN cd /usr/app && git clone https://github.com/santacodes/SecureEx.git

RUN cd /usr/app/SecureEx/server && go build

EXPOSE 3000

CMD [ "/usr/app/SecureEx/server/server" ]
