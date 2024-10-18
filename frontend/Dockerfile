# React frontend
FROM node:18 AS frontend

WORKDIR /todo-react

# Copy package.json and yarn.lock
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install

# Copy the rest of the application code
COPY . .

# Build the application
RUN yarn build

# Nginx stage to serve the build
FROM nginx:alpine

# Copy the build files to the nginx server
COPY --from=frontend /todo-react/dist /usr/share/nginx/html

# Expose the default port for nginx
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
