FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY main.html /usr/share/nginx/html/
COPY style.css /usr/share/nginx/html/

EXPOSE 80
