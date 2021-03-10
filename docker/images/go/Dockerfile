FROM golang:1.15.5-buster


RUN apt-get -y update --fix-missing
RUN apt-get upgrade -y

RUN apt-get -y install \
	apt-utils \
	net-tools \
	openssh-server \
	autossh \
	wget \
	nano \
    git \
	curl \
    openssl \
	unzip \
    zip \
    apt-file

    WORKDIR /app

    RUN go get github.com/githubnemo/CompileDaemon
    ENTRYPOINT CompileDaemon --build="go build -o /go/bin ./main.go" --command="/go/bin/main"