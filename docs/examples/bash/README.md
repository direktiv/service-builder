# Bash Service

# Init

```
docker run --user 1000:1000 -v `pwd`:/tmp/app direktiv/service-builder init bash-service
```

# Service

Because this service will be implemented in `bash`, we'll utilize three commands to count word occurrences on a webpage with curl. The first command is curl to get actually get it, and the other ones are grep and wc to do the actual searching and counting. But, before we go any further, let's take a look at the input.

# Input

We need to accept two arguments. An `address` and a `search` phrase. Both are string values, and they are required to make this service operate. We need to edit the body input configuration to read as follows:


```yaml
- name: body
  in: body
  schema:
    type: object
    required:
      - address
      - search
    properties:
      address:
        type: string
      search:
        type: string
```

# Command

To achieve this functionality we use `curl` to get the webpage and `grep` the search string. At the end we are word counting (`wc`) the results. That translates into this command:

```yaml
exec: |-
    bash -c 'curl -sL {{ .Address }} | grep -o -i {{ .Search }} | wc -l'
```

As you can see we are using [go templating](https://pkg.go.dev/text/template) here. Because how go templating works we need to upper-case the variables. This template means we are replacing `{{ .Address }}` with the incoming value as well as `search` in the `grep` instruction. For now don't want to change the output to something custom so we can comment that out for now. The `x-direktiv` section should look like this now:

```yaml
x-direktiv:   
    cmds:
    - action: exec
        exec: |-
        bash -c 'curl -sL {{ .Address }} | grep -c {{ .Search }}'
    # output: |
    #   {
    #     "greeting": "{{ index (index . 0) "result" }}"
    #   }
```

The code is now ready to be generated and run. The procedure for generating the code is straightforward, and it only takes one command to execute in the folder where the init command was run.

```
docker run --user 1000:1000 -v `pwd`:/tmp/app direktiv/service-builder gen v1.0.0
```

Now we can start the service with the `run.sh` file in the `v1.0.0` folder. This should compile the application and start the service on port 8080. To test it we can either use curl or call the UI of the service at http://127.0.0.1:8080/docs.

```
curl --header "Content-Type: application/json" \
  --header "Direktiv-ActionID: development" \
  --request POST \
  --data '{"address":"http://www.direktiv.io","search":"meta"}' \
  http://localhost:8080
```

This execution will fail and in the logs of the running service container we can see the error message:

```
bash: line 1: curl: command not found
```

This can be fixed by changing the generated Dockerfile to add the installation of `curl` and `grep`.


```Dockerfile
...
FROM ubuntu:21.04

RUN apt-get update && apt-get install ca-certificates -y

# make sure curl and grep are installed
RUN apt-get update && apt-get install curl grep -y

COPY --from=build /application /bin/application

EXPOSE 8080

CMD ["/bin/application", "--port=8080", "--host=0.0.0.0"]
```

If we run the service again with `run.sh` and executing a search again it will return with a successful response.

```json
[
	{
		"result": 6,
		"success": true
	}
]
```

Our service is working but we want to make the output a little bit more readable. For this we can change the `output` we commented out before and change the name. The output is using a template as well and it is actually quite simple. Without an output specified the service returns an array of the results of the commands it has executed. In this case it was only one. To get the first item in the list we use `(index . 0)` and from this first item we want the `result` value: `{{ index (index . 0) "result" }}`. The final swagger.yaml file can be seen [here](v.1.0.0/swagger.yaml).

```yaml
x-direktiv:   
    cmds:
    - action: exec
        exec: |-
        bash -c 'curl -sL {{ .Address }} | grep -c {{ .Search }}'
    output: |
      {
        "hits": "{{ index (index . 0) "result" }}"
      }
```

After every change in the command the service needs to be recreated. After that it is ready to be published and used in Direktiv. 

```yaml
functions:
- id: search
  image: localhost:5000/websearch
  type: knative-workflow
states:
- id: getter 
  type: action
  action:
    function: search
    input: 
      address: http://www.direktiv.io
      search: jq(.search)
```