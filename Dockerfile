FROM golang:1.14.0-alpine AS module-init
WORKDIR /app
COPY . .
RUN go mod download 
#Building Application
RUN CGO_ENABLED=0 GOOS=linux go build -o scrmabled-strings .
#Running Application
FROM busybox:stable
WORKDIR /app
#Copy the dependent files
COPY --from=module-init /app/dict.txt .
COPY --from=module-init /app/strings.txt .
COPY --from=module-init /app/scrmabled-strings .
#Run the Binary
ENTRYPOINT ["/app/scrmabled-strings","--dictionary=/app/dict.txt","--input=/app/strings.txt"]