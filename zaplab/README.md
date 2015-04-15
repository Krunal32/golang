Run instructions: 

cd to zapserver folder run "go build"
part1
task a)   ./zapserver -lab=a
task c1)  ./zapserver -lab=c1
task c2)  ./zapserver -lab=c2
task e)  ./zapserver -lab=e 

part2a)

run rpc server with: ./zapserver -lab=f
run client: cd to folder "rpc_client" 
go run client.go -subscription<top10 or duration> -delay=<number> -address=<ip address>
client will connect on a random accessible port  on localhost
rpc server  port is on 12110
