##Interacting with the REST API using Curl

# Assumptions
    Takes in JSON binding
    Localhost:9090 is enabled OR 127.0.0.1 (loopback)


# POST
    curl -X POST localhost:9090/todos -d '{"id":"1", "item":"Read novel","completed":false}'

# GET
    curl localhost:9090/todos
