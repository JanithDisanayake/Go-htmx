#!/bin/bash

# Check if figlet is installed
if ! command -v figlet &> /dev/null; then
    echo "Figlet is not installed. Please install figlet before running this script."
    exit 1
fi

# Print "Janix" in large letters using figlet
figlet "Janx"
