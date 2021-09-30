# go-Seidon
Considered as one of the king of the storage. With help of the others king around the ocean, He is manage your application files spread around the internet.

![go-Seidon logo][goseidon-logo]

## 🚀 Motivations
1. Provide single point of entry to upload file.
2. Provide single point of entry to accessing the uploaded file.
3. Allowing multiple files upload at once.
4. Allowing multiple storage `provider`, current support is: `Built in` or `AWS S3`
5. Support multiple `language` interface, currently supported: `id`, `en`

## 📋 Documentation
- [**API Documentation**](doc/API.md)
- [**Business Documentation**](doc/BUSINESS.md)
- [**Database Documentation**](doc/DATABASE.md)
- [**Testing Documentation**](doc/TESTING.md)
- [**Relevant Document**](doc/DOCUMENT.md)

## 👷🏻 Architecture
![System Architecture][architecture-image]

## 🖖 Dependencies
### Inbound Dependencies
1. Client with a given `application` configuration (app_key, app_secret)
2. Public resource through `/file/:id` endpoint
3. Database (MongoDB)

### Outbound Dependencies
1. Built in file service
2. AWS S3
3. Database (MongoDB)

## ✔️ Running the App

### Setup
- Copy `.env.example` into `.env`
- Setup `.env` configuration

### Development
```
$ go run cmd\web-service\main.go
```

### Deployment

Adjust deployment according to production pipeline, e.g: using `docker`, etc.
But in general we can simply run with the following command:

```
$ go run cmd\web-service\main.go
```

## ❓ QnA

No QnA right now

## 👀 Known Issues

No issues right now

## 💪 Todo
1. Validation
2. Local GetDetail implementation
3. AWSS3 SaveFile & GetDetail implementation
4. BuiltIn SaveFile & GetDetail implementation
5. Database model + repository implementation
6. GetDetail API
7. GetResource API
8. Translation
9. Uni Test
10. End to end Test

## 🤩 Nice to Have
1. Depedency injection
2. Allowing file authorization in the future.
3. Custom file slug configuration.
4. Caching layer
5. Custom file validation rules
6. Storage dashboard monitoring
7. Multiple provider for each file support (for backup purpose)

## 💖 Contributions

Please always follow the development guidance described above to keep the `code quality` great and also to `decrease unecessary bug`. 

Do run the `unit test` after changing code or before push/updating the code.

> *Leave better than you found it*

[goseidon-logo]: asset/image/go-seidon.png?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
[architecture-image]: asset/image/system-architecture.jpg?raw=true
