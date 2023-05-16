# isolation levels and related issues

## Read Phenomena
### Dirty Read 
    * caused by add or updated transaction which reads wrong data and this might show wrong data and cause issues when rolling back
    * A transaction reads data written by other concurrent uncommitted transactions
  
### Non-Repeatable Read
    * A transaction reads the same row twice and sees different value because it has been modified and by other committed transaction

### Phantom Read
    * A transaction re-executes the same query and sees different set of rows because database has been modified and by other committed transaction

### Serialization Anomaly
    * The result of a group of concurrent committed transactions is impossible to achieve if we try to run them sequentially in any order without overlapping

# Standard isolation Levels
* Defined by ANSI- American National Standard Institute
*  Read-Uncommitted --> See data written by uncommitted transactions
*  Read-Committed --> only See data written by committed transactions --> This means dirty read condition is no longer possible
*  Repeatable Read --> Same read query will return same result no matter how many times we execute a query even when new transactions have committed new changes that satisfy the query
*  Serializable -->Concurrent transactions that are running at this level are guaranteed to be able to yield same result as if they are executed sequentially in some order without overlapping


# Reading and configuring variables from environment variables
## "viper" and its advantages when it comes to reading environment variables

1. It can Find,Load, Unmarshal config file - JSON,TOML,YAML,ENV,INI
2. Override existing values, set default values
3. Read config from remote system like etcd,Consul etc
4. Live watching and writing config file