# go-Seidon
With help of the other `kings` around the `ocean`, He is considered as one of the king of storage managers. 
He can manage your application files which spread over the internet.

![go-Seidon logo][goseidon-logo]

## 🚀 Motivations
1. Provide single point of entry to upload file.
2. Provide single point of entry to accessing the uploaded file.
3. Allowing multiple files upload at once.
4. Allowing multiple storage `provider`, current support is: `local`
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
```bash
# run using hot reloading #
$ air 
$ air -c .air.toml # custom configuration

# run manually #
$ go run \\cmd\\web-service\\main.go 
```

### Deployment

Adjust deployment according to production pipeline, e.g: using `docker`.
But in general we can simply run with the following command:

```bash
# build manually #
$ go build -o ./build/web-service/ ./cmd/web-service/main.go

```

## ❓ QnA

No QnA right now

## 👀 Known Issues

No issues right now

## 💪 Todo
1. Don't need to store the `file.type` (enough with extension and mimetype)
2. Refactor file entity used amoung package (it's better to have little copying rather than little dependency)
3. Refactor file size rule to be adjusted according to file extension/mimetype
4. Local GetDetail implementation
5. AWSS3 SaveFile & GetDetail implementation
6. BuiltIn SaveFile & GetDetail implementation
7. Database model + repository implementation
8. GetDetail API
9. GetResource API
10. Translation
11. Unit Test
12. End to end Test

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
