# Build the Go server
FROM golang AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o go_server

# Build the Svelte app
FROM node:21-alpine AS node-builder
WORKDIR /app
COPY package.json package-lock.json postcss.config.js   ./
COPY svelte.config.js tsconfig.json   ./
COPY tailwind.config.js vite.config.ts   ./
COPY src/ static/ ./
RUN npm install
COPY . .
# RUN npm run build

# Final stage
FROM node:21-alpine
WORKDIR /app
COPY --from=go-builder /app/. /app/.
COPY --from=node-builder /app/. .
COPY ./entrypoint.sh /app/
RUN chmod +x /app/entrypoint.sh

EXPOSE 4696 
EXPOSE 5173
CMD ["/app/entrypoint.sh"]