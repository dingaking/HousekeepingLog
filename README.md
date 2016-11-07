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

# update api admin
cd {workspace}<br>
apidoc -o output<br>
ex) cd ~/Documents/workspace/HousekeepingLog/src/apidoc_admin;<br>
apidoc -o output

# update api user
cd {workspace}<br>
apidoc -o output<br>
ex) cd ~/Documents/workspace/HousekeepingLog/src/apidoc_user;<br>
apidoc -o output
