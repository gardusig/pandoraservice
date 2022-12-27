# grpc-service

The goal of this project is to document the setup for establishing a gRPC communication between a server and its clients.

## GuessNumberService

- This service randomly generates a number on its startup and refreshes it whenever someone guesses it right
- Guessed numbers should lie within this interval: $[-10^{18}, +10^{18}]$
- The prize is a locked box
- You can unlock it with the server and see its value

There are two defined procedures enabling clients to remotely call it:
- `GuessRandomNumber`
  - Input: `Guess`
  - Output: `GuessResponse`
- `OpenBox`
  - Input: `LockedBox`
  - Output: `Box`

[Check it out](/protos/example.proto)!

## Languages

1. [Go](/go/README.md)
