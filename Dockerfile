FROM alpine:latest

ENV TZ=Asia/Shanghai

WORKDIR /app

# Copy executable from Jenkins build
COPY main /app/main

EXPOSE 8080

# Set executable permission and run
RUN chmod +x /app/main
CMD ["./main"]
