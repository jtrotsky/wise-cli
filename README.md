## wise-cli

### why
Help you integrate Wise into your application.

- get and display quotes 
- view / transfer between balances 
- add / update / delete recipients 
- test webhooks from the cli

### usage
```
play with the Wise API

Usage:
 wise quote create --api-key <YOUR-API-SECRET-KEY-HERE> --to NZD --from GBP --sourceAmount 100

 1.91 ┤╭╮   ╭╮╭─╮      ╭╮         ╭─╮
 1.90 ┼╯╰╮╭─╯╰╯ ╰─╮ ╭──╯│     ╭─╮╭╯ ╰
 1.89 ┤  ╰╯       ╰─╯   ╰╮   ╭╯ ╰╯
 1.88 ┤                  ╰╮  │
 1.87 ┤                   ╰─╮│
 1.86 ┤                     ╰╯
                   30 days

 Quote for 100 GBP to NZD at 1=1.899290
  -> 188.26 NZD will arrive in 57h

Flags:
  -h, --help   help for using wise

Additional help topics:
  wise quote  List quote commands
  wise transfer List transfer commands
  wise recipient List recipient commands
  wise balance List balance commands
```
