### BUILD ###
FROM node:16-alpine AS BUILD
WORKDIR /tmp/app
COPY . .
RUN npm install && npm run build
### RUN ###
FROM nginx:alpine
COPY --from=BUILD /tmp/app/dist/angular-app /usr/share/nginx/html
