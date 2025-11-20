FROM node:25 AS step1
WORKDIR /myapp
COPY package*.json ./

# RUN npm config set registry "http://host.docker.internal:8081/repository/demo01"
RUN npm --loglevel info install
COPY . .
RUN npm run build
RUN ls /myapp/
# /myapp/dist

FROM nginx:1.29.3
COPY --from=step1 /myapp/dist /usr/share/nginx/html