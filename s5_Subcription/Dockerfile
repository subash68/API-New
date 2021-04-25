# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /ch-subscription

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o subscription

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./subscription" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /ch-subscription-test

# Copy source files from dev
COPY --from=dev /ch-subscription/ /ch-subscription-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /ch-subscription-prod

# Coping only Executable file from dev
COPY --from=dev /ch-subscription/subscription /ch-subscription-prod/

# Start command
CMD ["./subscription"]