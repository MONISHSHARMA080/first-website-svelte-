# Build the Go server
FROM golang AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -x -o go_server

# Build the Svelte app
FROM node:21-alpine AS node-builder
WORKDIR /app
COPY package.json package-lock.json postcss.config.js svelte.config.js  tsconfig.json   ./
RUN npm install
# COPY . .
COPY ./entrypoint.sh /app/
COPY src/ /app/src/.
RUN chmod +x /app/entrypoint.sh
# RUN npm run build

# Final stage
# FROM node:21-alpine
# WORKDIR /app
COPY --from=go-builder /app/. /.
# COPY --from=node-builder /app/. /.

EXPOSE 4696 5173
ENTRYPOINT ["/app/entrypoint.sh"]