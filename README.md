# simple-client-server-go

Start Server:
    go run server/server.go localhost:18081
Start Client 1:
    go run client/client.go localhost:28081 localhost:18081
Start Client 2:
    go run client/client.go localhost:28082 localhost:18081

Example message to Client 1:
    {"destination":"localhost:28082","transaction_id":1,"policy_id":2,"vm_address":"123"}
Example message to Client 2:
    {"destination":"localhost:28081","transaction_id":1,"policy_id":2,"vm_address":"123"}