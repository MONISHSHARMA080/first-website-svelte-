FROM golang:1.22 as go-builder
WORKDIR /app
COPY a.go ./
COPY go.sum ./
COPY go.mod  ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -x -o go_server

FROM node:21-alpine as build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json tailwind.config.js  vite.config.ts postcss.config.js svelte.config.js  tsconfig.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY src/ ./src/
COPY static/ ./static/

COPY --from=go-builder /app/go_server /app/go_server
COPY ./entrypoint.sh /app/
RUN chmod +x /app/entrypoint.sh
# Expose the port that the app runs on
EXPOSE 5173 4696

# Command to run the app in development mode
ENTRYPOINT ["/app/entrypoint.sh"]
# docker build -t aa .; docker run -p 4696:4696 -p 80:5173 aa