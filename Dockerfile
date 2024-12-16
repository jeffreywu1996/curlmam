FROM golang:1.20

WORKDIR /app

# Install necessary dependencies
RUN go install github.com/cosmtrek/air@latest

# Copy application code
COPY . .

# Expose the port used by the app
EXPOSE 8080

# Run the Go app
CMD ["air", "-c", ".air.toml"]
