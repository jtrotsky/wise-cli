> **Unmaintained.** The last functional change here was March 2023. It has not been
> tested against the Wise API since, the `recipient` and `webhook` commands were never
> finished, and its dependencies (Go 1.16, cobra 1.1.3, viper 1.7.1) are years out of
> support. Treat it as a read-only example of calling the Wise API from Go, not as a
> working tool. Not affiliated with or endorsed by Wise.
>
> Note that the example below passes an API token as a command-line argument, which
> leaves it in shell history and the process table. Read it from the environment instead.

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
