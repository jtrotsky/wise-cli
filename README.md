## wise-cli

### why
- get and display quotes 
- view / transfer between balances 
- add / update / delete recipients 
- test webhooks from the cli

### usage
```
-k --api-key Your secret API key 
```

### example
```
./wise-cli login --api-key <YOUR-API-SECRET-KEY-HERE> 
./wise-cli listen
./wise-cli trigger balance.deposit
```
