# Multi-User Chat System (MUCS)
 [中文版](README.zh.md)

Multi-User Chat System (MUCS) is a simple instant messaging solution that allows users to register, log in, and join different chat rooms for communication through a client application. This project is developed in Go language, demonstrating an implementation method for a server/client architecture based on TCP.
![](/img/mes.png)

## Features
- User Registration and Login: New users can register an account and log in to the system with that account.
- Multi-User Chat Rooms: Users can join public chat rooms to communicate with other online users.
- Real-time Message Exchange: Supports the instant transmission of text messages.
- TCP-based Network Communication: Builds stable client and server communication using Go language's network programming interface.
- Data Storage: Utilizes Redis for storing and managing user information and chat history.

## Getting Started
The following guide will help you install and run this project on your local machine for development and testing purposes.

- Go 1.19.3
- Redis Server

### Installation
1. Clone the repository:
```
git clone https://github.com/ZhijiunY/MUCS.git
```
2. Start the Redis Server
Ensure your Redis server is running.
![](/img/redis.png)

3. Run the Server
Execute the following command in the server directory of the project:
```
./server.exe
```
![](/img/server.png)

4. Run the Client
In another terminal window (open two of these), run the client application:
```
./client.exe
```
![](/img/client.png)

## Usage Instructions
- After the client application starts, follow the prompts to log in, sign up, or exit.
- Once logged in or registered, you will be able to join chat rooms and communicate with other users.

## Architecture
This project adopts a client-server model, mainly consisting of two parts: the client application and the server application. The server is responsible for handling user requests, message forwarding, and data persistence. The client provides a user interface, supporting user interactions.

## Contribution
Contributions of any kind are welcome. For significant changes, please open an issue first to discuss what you would like to change.

## Version 
0.1