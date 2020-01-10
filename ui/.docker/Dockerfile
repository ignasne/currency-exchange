FROM node:12.13-alpine as builder

RUN mkdir -p /opt/source
WORKDIR /opt/source

COPY public/ /opt/source/public
COPY src/ /opt/source/src
COPY package.json /opt/source
COPY yarn.lock /opt/source

WORKDIR /opt/source

RUN yarn
RUN yarn build

FROM nginx:1.12-alpine

COPY --from=builder /opt/source/build /usr/share/nginx/html
COPY .docker/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD /bin/sh -c "nginx -g 'daemon off;'"
