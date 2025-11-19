FROM node:25 AS step1
WORKDIR /myapp
COPY . .
RUN npm install
RUN npm run build
# /myapp/dist

FROM nginx:1.29.3
COPY --from=step1 /myapp/dist /usr/share/nginx/html