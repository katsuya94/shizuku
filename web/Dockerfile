FROM node

RUN mkdir /usr/local/web
WORKDIR /usr/local/web

ADD . .

ENTRYPOINT ["/usr/local/web/entrypoint.sh"]
CMD ["npx", "webpack-dev-server", "--mode", "development"]
