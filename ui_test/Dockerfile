FROM mcr.microsoft.com/playwright:v1.56.1-noble
WORKDIR /app
COPY . .
RUN npm install
RUN npx playwright install
CMD [ "npx", "playwright", "test" ]