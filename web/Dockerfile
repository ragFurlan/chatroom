FROM node:16 AS build
WORKDIR /app
COPY package.json /app
RUN npm i npm@latest -g
RUN npm install

RUN npm install -g @angular/cli
COPY ./ /app
RUN npm run build

FROM nginx:alpine
RUN rm -rf /usr/share/nginx/html/*
COPY nginx.conf /etc/nginx/conf.d/custom.conf
COPY --from=build /app/dist/web /usr/share/nginx/html


EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]


# FROM node:16 AS build
# WORKDIR /appd
# COPY package.json /app
# RUN npm i npm@latest -g
# RUN npm install
# COPY ./ /app

# RUN ng build --configuration=production   

# RUN npm run build

# # Segunda etapa: Criar imagem final com o build otimizado
# FROM nginx:alpine
# RUN rm -rf /usr/share/nginx/html/*
# COPY nginx.conf /etc/nginx/conf.d/custom.conf
# COPY --from=build /app/dist/web /usr/share/nginx/html

# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]

# Primeira etapa: Build do projeto e otimização do CSS
# FROM node:16 AS build
# WORKDIR /app
# COPY package.json /app
# RUN npm i npm@latest -g
# RUN npm install

# # Instale o Angular CLI
# RUN npm install -g @angular/cli

# COPY ./ /app

# # Otimizar o build do Angular
# RUN ng build --configuration=production

# # Segunda etapa: Criar imagem final com o build otimizado
# FROM nginx:alpine
# RUN rm -rf /usr/share/nginx/html/*
# COPY nginx.conf /etc/nginx/conf.d/custom.conf
# COPY --from=build /app/dist/web /usr/share/nginx/html

# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]
