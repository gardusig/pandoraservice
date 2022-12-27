# grpc-service

The goal of this project is to document the setup for establishing a gRPC communication between a server and its clients.

## GuessNumber Service

### Statements

- A special number is randomly generated at the server startup
  - It gets updated every time a guess is made by a client with the exact same number
- For each guess, the server will respond with
  - `<` if given number is less than special number
  - `>` if given number is greater than special number
  - `=` if given number is equal to special number
    - There will also be a locked box as your prize
- Guessed numbers must lie within the interval defined by: $[-8 * 10^{18}, +8 * 10^{18}]$
  - Given that your computer probably can't process more than $10^{10}$ operations in a second, a naive approach of guessing each possible number would take at least $5$ years to execute
- There is in fact one algorithm capable of solving this question a bit faster - `binary search`
  - Instead of guessing each number, you can benefit from the server response to avoid making unnecessary guesses
  - Since each wrong guess reduces the amount of possibilities by its half
    - Worst case goes from $O(N)$ to $O(log_2(N))$
      - where $N$ is the amount of possible values ($\approxeq16 * 10^{18}$)
        - which also happens to be quite close to the amount of possible numbers to fit inside a 64-bit integer: $log_2(16 * 10^{18})\approxeq64$

### Protocol Buffers

There are two defined procedures to be remotely called
- `GuessRandomNumber`
  - Input: `Guess`
  - Output: `GuessResponse`
- `OpenBox`
  - Input: `LockedBox`
  - Output: `Box`

[Check it out](/protos/example.proto)!

## Programming Languages

1. [Go](/go/README.md)
