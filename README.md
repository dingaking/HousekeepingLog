# HousekeepingLog GOPATH
cd {workspace}<br>
export GOPATH="\`pwd\`"<br>
ex) cd ~/Documents/workspace/HousekeepingLog;<br>
export GOPATH="\`pwd\`"

# HousekeepingLog Help
cd $GOPATH;<br>
godoc -http=:6060;

# start testing
cd $GOPATH/TESTS/;<br>
go test -v

# run database
cd {workspace}<br>
mongod --dbpath {db_path}<br>
ex) cd ~/Documents/workspace;<br>
mongod --dbpath ./MongoDB/ &

# run server
cd {workspace}<br>
go build; ./apis<br>
ex) cd ~/Documents/workspace/HousekeepingLog/src/apis;<br>
go build; ./apis
