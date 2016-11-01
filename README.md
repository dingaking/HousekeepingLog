# HousekeepingLog Help
cd $GOPATH;<br>
godoc -http=:6060;

# start testing
cd $GOPATH/TESTS/;<br>
go test -v

# run database
cd {workspace}<br>
mongod --dbpath {db_path}<br>
ex) cd ~/Document/workspace;<br>
mongod --dbpath ./MongoDB/ &

# run server
cd {workspace}<br>
go build; ./apis<br>
ex) cd ~/Document/workspace/HousekeepingLog/src/apis;<br>
go build; ./apis
