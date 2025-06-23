##Interacting with the REST API using Curl

#POST
curl -X POST localhost:9090/todos -d '{"id":"1", "item":"Read novel","completed":false}'
