# frontend/Dockerfile
FROM node:20-alpine

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Serve with built-in static server
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", "build"]
