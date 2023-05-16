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

## Mocking database for running tests instead of using normal database

### Independent Tests
* Independent tests are better handled with mock dtabase because they are writingand reading data from memory and not a database so tests dont need to worry about other tests hogging each other.

### Faster Tests
* Since the mimic database is in memory, it reduces a lot of time used intalking to database.

### Test Coverage 
* we can cover super edge cases like connection lost with database etc which are very very very rare in  real scenario.

'''
    mock database will work fine just like normal database because we have already tested tests in normal database till now and another thing is that this database is and the origianal database can and should have same interface thus reducing any issues that might happen bceause of improper configuration
'''
* Mock Databases can be implemented in 2 ways 
    1. Using a in memory store
    2. using stubs which mimic database