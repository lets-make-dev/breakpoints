#!/bin/bash

# Build the summarize executable
go build -o bp

# Make the twig executable
chmod +x bp

# Move to bin
sudo mv ./bp /usr/local/bin

