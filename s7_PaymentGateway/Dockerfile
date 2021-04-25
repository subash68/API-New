# Base image 
FROM golang:1.15-alpine AS dev

# Working directory
WORKDIR /ch-paymentgateway

# Coping All files to 
COPY . .

# Runing dependencies
RUN go get github.com/githubnemo/CompileDaemon && go build -o paymentgateway

# Exposing Port
EXPOSE 8080

# Command uses if not Overriden
CMD [ "./paymentgateway" ]

# building Image from dev
FROM dev AS qa

# Testing  working directory
WORKDIR /ch-paymentgateway-test

# Copy source files from dev
COPY --from=dev /ch-paymentgateway/ /ch-paymentgateway-test

# Command to test
CMD ["go test > test_report.log"]

# Building Production image
FROM dev as prod

# Production working directory
WORKDIR /ch-paymentgateway-prod

# Coping only Executable file from dev
COPY --from=dev /ch-paymentgateway/paymentgateway /ch-paymentgateway-prod/

# Start command
CMD ["./paymentgateway"]