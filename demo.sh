#!/bin/bash
set -e

echo "Starting server in background..."
go run main.go &
SERVER_PID=$!

# Wait for the server to start
sleep 3

echo "Running demo command..."
curl http://localhost:8080/people

echo ""
echo "Killing server..."
kill $SERVER_PID
