Problem: Log Processor

Part 1

Write a Go program simulating a multi client-server log processing system.

The clients send requests to the server specifying a desired log type (e.g., "error") and client id.

The program reads client requests from a JSON file clients.json

The server should process log entries from a JSON file logs.json using concurrency, with multiple threads each
responsible for processing 10 lines of the data. Each thread processes a consecutive set of 10 lines.
The first thread should process the first 10, the second thread should process the second 10, and so on.

The server counts how many entries match the requested log type and sends this aggregated count back to the client.

The clients receive and print the count.

Use channels for communication between client and server.

Ensure concurrent processing of the log file for performance.

Handle errors gracefully.

*This is a simulated system. You are not building a real socket- or HTTP-based server.
Instead, simulate the clients and server within a single Go program

Part 2

Extend the program to store recent log counts per log type in a concurrency-safe in-memory structure,
so they can be safely accessed later (e.g., by an external API).


Please check go file