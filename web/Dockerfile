FROM node:16 AS build
WORKDIR /app
COPY package.json /app
RUN npm i npm@latest -g
RUN npm install
COPY ./ /app
RUN npm run build

FROM nginx:alpine
RUN rm -rf /usr/share/nginx/html/*
COPY nginx.conf /etc/nginx/conf.d/custom.conf
COPY --from=build /app/dist/web /usr/share/nginx/html


EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
