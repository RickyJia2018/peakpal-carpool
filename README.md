# gRPC server built in go

### swagger url `:8080/swagger`



PostMan Script to get access token
```

for(const key in pm.response.messages){
    var temp = pm.response.messages[key];
    if (temp && temp[0].data){
        let data = temp[0].data
    console.log(data)
    pm.collectionVariables.set("access_token", "bearer "+data.access_token);
    break;
    }
}

```