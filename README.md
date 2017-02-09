masscan wrapper as a http service


 git clone  https://github.com/meixinyun/masscan-httpservice.git
 
 cd $GOPATH/src/github.com/meixinyun/masscan-httpservice
 
 go get ./...
 
 ./control build
 
 
 sudo ./jaguar-masscan-httpservice -c cfg.json
 
 
open exploer, 

http://127.0.0.1:20000/masscan?cmd=--range ip1-ip2  -p80


you will find the result as follows:
 
 {   "ip": "61.135.169.125",   "timestamp": "1486643473", "ports": [ {"port": 80, "proto": "tcp", "status": "open", "reason": "syn-ack", "ttl": 55} ] },
{finished: 1}
