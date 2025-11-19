FROM node:25 AS step1
WORKDIR /myapp
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
RUN ls /myapp/
# /myapp/dist

FROM nginx:1.29.3
COPY --from=step1 /myapp/dist /usr/share/nginx/html