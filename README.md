## Summary

The goal of this project is to document the setup for establishing a gRPC communication between a server and its clients. Execution details at GitHub Actions.

## Statements

Guess the right number and receive a $PandoraBox$.

![Wuya](pictures/wuya.webp) ![PandoraBox](pictures/box.webp)

Beware that it's initially locked and you must ask the server to open it for you :)

![OpenedBox](pictures/opened.jpg)

### Rules

- A special number is randomly generated at the server startup
  - Guessed numbers must lie within the interval: $[-4 * 10^{18}, +4 * 10^{18}]$
  - It gets updated every time a guess is correctly made by a client
- For each guess at `GuessNumber` $rpc$, the server will respond with:
  - $<$ if given number is less than special number
  - $>$ if given number is greater than special number
  - $=$ if given number is equal to special number
- A right guess will return a $LockedPandoraBox$ object
  - It contains an encrypted message
- You can ask the server to open it at `OpenBox` $rpc$
  - The response will have an $OpenedPandoraBox$ object with a decrypted message

### Observations

- Given that your computer probably can't process more than $10^{10}$ operations in a second, a naive approach of guessing each possible number would take at least $25$ years to execute
  - Worst case: $8 * 10^{18}$ operations
  - Operations per second: $10^{10}$
  - Total in seconds: $8*10^{8}$
  - Total in years: $25$
- There is in fact one algorithm capable of solving this question a bit faster - $BinarySearch$
  - Instead of guessing each number, you can benefit from the server response to avoid making unnecessary guesses
  - The server response always tell you if the answer is less or greater. If it's less, making guesses with greater numbers does not help. Neither if it's greater would make sense to make a guess with lesser numbers.
  - So each wrong guess reduces the amount of possibilities by its half
    - Worst case goes from $O(N)$ to $O(log_2(N))$

### Protocol Buffers

Module created at separate repository. [Check it out](https://github.com/gardusig/pandoraproto)!

#### Running

```bash
docker build -t grpc_service --progress=plain .
```
