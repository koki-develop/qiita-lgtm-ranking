FROM golang:1.16
WORKDIR /var/task
ENV HOME /root

# install nodejs and yarn
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash - \
 && apt install -y nodejs \
 && npm install --global yarn
