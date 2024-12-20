# telnet chat server
this is a pretty simple chat server
if you want to try it run the tcp_based_chat.exe and connect the server using telnet/putty


# telnet chat server
If you want to try it, run the [tcp_based_chat.exe](./tcp_based_chat.exe) and connect to the server using telnet/putty.

### commands
`/nick` pick your nickname (if you don't do that you shall be named anonymous)

`/join` join/create a chat room

`/rooms` list existing chat rooms

`/msg` use this to write messages to your pals

`/quit` leave

### config
configuration is quite simple
all you need to do is write the port you want to run the server on in [config.json](./config.json) in the following format `"port": ":<you port number>"`
then run the server and you're good