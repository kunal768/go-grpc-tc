# go-grpc-tc
Go gRPC server for performing unary gRPC requests

### To Build

```shell
go build -o go-grpc-tc
```

### To Run 

```shell
go run main.go
```

### Run Unit Tests 

```shell
go test ./user    
```

### For Accessing gRPC endpopints :
### Use a gRPC client tool like POSTMAN and upload the [proto](./proto/userservice.proto) file 
> [!IMPORTANT]  
> After uploading the proto file Set url to "localhost:8080" and select appropriate service from dropdown at the right <br />
> Example screenshots are attahced

> [!NOTE]  
> On starting the server two users are seeded by default into the databases, they are as follows : <br />
> 1: {ID: 1, FName: "John", City: "New York", Phone: 1234567890, Height: 180.5, Married: true} <br />
> 2: {ID: 2, FName: "Jane", City: "Los Angeles", Phone: 9876543210, Height: 165.2, Married: false}


#### Request/Response examples for gRPC server 

#### Get User By Id
##### Request 
```json
{
    "id": 1
}
```

##### Response 
```json
{
    "user": {
        "id": 1,
        "fname": "John",
        "city": "New York",
        "phone": "1234567890",
        "height": 180.5,
        "married": true
    }
}
```


#### Get User By Id (Invalid)
##### Request 
```json
{
    "id": 0
}
```

##### Response 
```json
{
  Invalid argument
  You have specified an invalid argument.
}
```

#### Get Users By Id
##### Request 
```json
{
    "ids": [2, 45]
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 2,
            "fname": "Jane",
            "city": "Los Angeles",
            "phone": "9876543210",
            "height": 165.2,
            "married": false
        },
        {
            "id": 45,
            "fname": "reprehenderit consectetur exercitation velit Ut",
            "city": "Lorem aliqua",
            "phone": "8442039398",
            "height": 48909313.409837335,
            "married": false
        }
    ]
}
```


#### Search Users By Fname 

##### Request 
```json
{
    "fname": "Jane"
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 2,
            "fname": "Jane",
            "city": "Los Angeles",
            "phone": "9876543210",
            "height": 165.2,
            "married": false
        }
    ]
}
```

#### Search Users By City

##### Request 
```json
{
    "city": "New York"
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 1,
            "fname": "John",
            "city": "New York",
            "phone": "1234567890",
            "height": 180.5,
            "married": true
        }
    ]
}
```

#### Search Users By Phone
##### Request 
```json
{
    "phone": "8442039398"
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 45,
            "fname": "reprehenderit consectetur exercitation velit Ut",
            "city": "Lorem aliqua",
            "phone": "8442039398",
            "height": 48909313.409837335,
            "married": false
        }
    ]
}
```

#### Search Users By Married Status
##### Request 
```json
{
    "married": false,
    "searchmarried": true
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 2,
            "fname": "Jane",
            "city": "Los Angeles",
            "phone": "9876543210",
            "height": 165.2,
            "married": false
        },
        {
            "id": 45,
            "fname": "reprehenderit consectetur exercitation velit Ut",
            "city": "Lorem aliqua",
            "phone": "8442039398",
            "height": 48909313.409837335,
            "married": false
        }
    ]
}
```

#### Add Users Valid 
##### Request 
```json
{
    "city": "Lorem aliqua",
    "fname": "reprehenderit consectetur exercitation velit Ut",
    "height": 48909313.409837335,
    "id": 45,
    "married": false,
    "phone": "8442039398"
}
```

##### Response 
```json
{
    "user": {
        "id": 45,
        "fname": "reprehenderit consectetur exercitation velit Ut",
        "city": "Lorem aliqua",
        "phone": "8442039398",
        "height": 48909313.409837335,
        "married": false
    }
}
```

#### Add Users Invalid 
##### Request 
```json
{
    
}
```
##### Response 
```json
{
  Invalid argument
  You have specified an invalid argument.
}
```

#### List Users 

> [!NOTE]  
> By default page is set to 1 and and pageSize is 10 (i.e. if request body is empty )

##### Request 
```json
{
    "page": 1,
    "pageSize": 1
}
```

##### Response 
```json
{
    "users": [
        {
            "id": 2,
            "fname": "Jane",
            "city": "Los Angeles",
            "phone": "9876543210",
            "height": 165.2,
            "married": false
        }
    ]
}
```



### Screenshots



 

<img width="922" alt="Screenshot 2024-07-08 at 5 22 43 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/f35178df-a561-4e6e-b8c9-7fff7228b58f">
<img width="922" alt="Screenshot 2024-07-08 at 5 22 53 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/5e8659de-68f0-4e47-a1eb-882341105dec">
<img width="922" alt="Screenshot 2024-07-08 at 5 23 10 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/be24a82e-ebaa-4b7b-adc4-86eb768b2277">
<img width="922" alt="Screenshot 2024-07-08 at 5 23 23 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/f5f784bc-94aa-40f9-abaa-04e2893298ee">
<img width="922" alt="Screenshot 2024-07-08 at 5 23 35 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/98971319-89db-47bf-a6c6-9c9111dc8e47">
<img width="922" alt="Screenshot 2024-07-08 at 5 23 56 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/d9da0c75-f918-47ee-8e2e-c6c7b8fdeac0">
<img width="922" alt="Screenshot 2024-07-08 at 5 24 05 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/fb568cb8-e08f-40fe-8996-f1ec113ab79d">
<img width="922" alt="Screenshot 2024-07-08 at 5 24 12 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/550011f4-a320-4e62-8dff-4df7690a51b2">
<img width="922" alt="Screenshot 2024-07-08 at 5 24 19 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/46f77f9c-2da9-4f3b-8ead-cf11e3da549f">
<img width="922" alt="Screenshot 2024-07-08 at 5 24 27 PM" src="https://github.com/kunal768/go-grpc-tc/assets/33108756/6dffb4e2-3ddc-4a5f-84d7-54ac3fcdd678">












