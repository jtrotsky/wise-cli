## wise-cli

### why
Use your Wise account from the command line and see an example of using 
Wise APIs with Go.
- see the rates and arrival time to send money anywhere 
- view and transfer funds between balances 
- add / update / delete recipients (TODO)
- test webhooks from the cli (TODO)

### usage
```
play with the Wise API

Usage:
 wise-cli quote create --to NZD --from GBP --amount 100 --token <YOUR-API-TOKEN-HERE> 

 1.91 ┤╭╮   ╭╮╭─╮      ╭╮         ╭─╮
 1.90 ┼╯╰╮╭─╯╰╯ ╰─╮ ╭──╯│     ╭─╮╭╯ ╰
 1.89 ┤  ╰╯       ╰─╯   ╰╮   ╭╯ ╰╯
 1.88 ┤                  ╰╮  │
 1.87 ┤                   ╰─╮│
 1.86 ┤                     ╰╯
                   30 days

 Quote for 100 GBP to NZD at 1=1.899290
  -> 188.26 NZD will arrive in 57h
```
