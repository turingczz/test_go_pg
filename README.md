# go-nft-status

These are steps to manually build a local go-nft-status.

## Fetch source code
We need to switch Makefile after fetch source code.
```bash
git clone https://gitlab.insidebt.net/btfs/go-nft-status.git
cd go-nft-status
```

## Install prerequisites  
Install GoLang on the machine.  
Install PostgreSQL on Mac and create the go-nft-status db run:
```bash
make install
```

## Config ENV for local instance of go-nft-status    
Replace username, password, host, db name for PostgreSQL database setting.
```bash
export DB_USER=<username>:<password>
export DB_HOSTNAME=localhost
export DB_NAME=db_nft
```

## run locally
```bash
make -f Makefile.local
```

## Build
```bash
make build
```
go-nft-status binary will be built into ./bin directory.

## Start the go-nft-server at the backend
Issue the following command:
```bash
cd bin
nohup go-nft-status </dev/null >/dev/null 2>&1 &
```
If copy and run go-nft-status in another directory, make sure to copy the IP2LOCATION DB along with go-nft-status binary.
