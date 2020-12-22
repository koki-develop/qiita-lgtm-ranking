FROM lambci/lambda:build-go1.x
WORKDIR /var/task
ENV HOME /root

# install nodejs
RUN curl -sL https://rpm.nodesource.com/setup_14.x | bash - \
 && yum -y install nodejs

# install yarn
ENV PATH $HOME/.yarn/bin:$HOME/.config/yarn/global/node_modules/.bin:$PATH
RUN curl -o- -L https://yarnpkg.com/install.sh | bash -s -- --version 1.22.10
