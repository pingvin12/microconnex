# RPC Service for Date Calculation

This is an RPC (Remote Procedure Call) service written in Go that provides a method to calculate the expiration date given a starting date and turnaround time.
## Description

The RPC service exposes a method called GetExpirationDate that takes a date in string format and a turnaround time in hours as input. It then calculates the expiration date by adding the turnaround time to the given date, considering working hours and skipping weekends. The result is returned in string format.
## Getting Started
### Prerequisites

To run this RPC service, you need to have Go installed on your system. If you don't have Go installed, you can download and install it from the official Go website: https://golang.org/
### Installation

Clone the repository to your local machine:



    git clone https://github.com/your-username/rpc-date-service.git

Navigate to the project directory:

    cd rpc-date-service

Build the Go executable:

    go build

### Running the Service

To start the RPC service, run the generated executable:


    ./rpc-date-service

By default, the service will listen on port 8080. You can change the port by modifying the main function in the main.go file.