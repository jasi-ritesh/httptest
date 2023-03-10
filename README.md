# httptest
This project intends to showcase the various testing strategies that can be used for 
performing blackbox testing.

## Expression Engine
```
This is the system under test and provides REST API's to interact with the Engine.
Following are the REST API's exposed by the system
1. Add Engine
Using this API one can create a new instance of the Engine
/engine?engine=name  --POST call at the URL and passing name of the eninge as query 
parameter

2. Delete Engine
Using this API to delete and instance of the Engine
/engine?engine=name -- Delete Call at the URL to delete the engine

3. Add Expression
Using this API one can add named expression to the Engine.
Multiple Calls can be made to add multiple expressions to the Engine

/engine/expr?engine=name -- POST Call at the URL
{"name":"First","expr":"2+3"}  --JSON data to be sent in body.
(JSON is an open standard file format and data intrchange 
format that uses human readable text to store and transmit data objects consisting
of attribute-value pairs and arrays.)

2. Evaluate
This API can be used to evaluate all the expressions currently in the Engine
/engine/evaluate?engine=name -- POST Call at the URL


3. Result
This API can be used to fetch all the expressions that have been evaluated by the engine.
It's returned as map
/engine/result?engine=name -- Get Call at the URL
result = {"First":"5"}

4. Delete Expression
This API can be used to delete a named expression from the Engine.
/engine/expr?name=First&engine=name -- Delete Call at the URL

5. Clear
This API can be used to Clear all named expression from the Engine
/engine/expr/clear?engine=name -- POST Call at the URL

```


## Running the Program and test cases
```
cd httptest
go run main.go

go clean --testcache & go test -v ./...

```


