FROM node:23-slim as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY  nginx/default.conf /etc/nginx/conf.d/default.conf
# COPY nginx/genmilk.crt /etc/nginx/certs/genmilk.crt
# COPY nginx/genmilk.key /etc/nginx/certs/genmilk.key
EXPOSE 80 433
CMD ["nginx", "-g", "daemon off;"]