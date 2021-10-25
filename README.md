# go-Seidon
With help of the other `kings` around the `ocean`, He is considered as one of the king of storage managers. 
He can manage your application files which spread over the internet.

![go-Seidon logo][goseidon-logo]

## 🚀 Motivations
1. Provide single point of entry to upload files.
2. Provide single point of entry to access the uploaded file.
3. Allowing multiple files upload at once.
4. Allowing multiple storage `provider`, current support is: `local`
5. Support multiple `language` interface, current supports are: `id`, `en`
6. Avoid coupling between storage service provider.

## 📋 Documentation
- [**API Documentation**](doc/API.md)
- [**Business Documentation**](doc/BUSINESS.md)
- [**Database Documentation**](doc/DATABASE.md)
- [**Testing Documentation**](doc/TESTING.md)
- [**Relevant Document**](doc/DOCUMENT.md)

## 👷🏻 Architecture
![System Architecture][architecture-image]

## ✔️ Running the App

### Setup
- Copy `.env.example` into `.env`
- Setup `.env` configuration

| Environment | Type | Example | Default Value | Description |
| --- | --- | --- | --- | --- |
| APP_URL | String | http://storage.domain.tld | (none) | Public application domain/subdomain used to access `goseidon` app |
| APP_HOST | String | localhost | localhost | Private application host used to access `goseidon` app privately, for example when used behind `load balancer` or `process management` |
| APP_PORT | Integer | 3000 | 3000 | Private application port used to access goseidon app privately |
| APP_DEFAULT_LOCALE | String | id | en | Default application langauge when no `Accept-Language` header or `lang` query specified |
| MIN_UPLOADED_FILE | Integer | 1 | 1 | Minimum amount of file to be uploaded in one single upload |
| MAX_UPLOADED_FILE | Integer | 5 | 5 | Maximum amount of file to be uploaded in one single upload |
| MIN_FILE_SIZE | Integer | 1 | 1 | Minimum file size `byte` for each uploaded file during single upload, default is 1 indicating valid `non zero` file size |
| MAX_FILE_SIZE | Integer | 134217728 | 134217728 | Maximum file size `byte` for each uploaded file during single upload, default is `134217728` byte or `128` MB |

### Development
```bash
# run using hot reloading #
$ air 
$ air -c .air.toml # custom configuration

# run manually #
$ go run \\cmd\\web-service\\main.go 

# testing #
$ ginkgo watch -r -v

```

### Deployment

Adjust deployment according to production pipeline, e.g: using `docker`.
But in general we can simply run with the following command:

```bash
# build manually #
$ go build -o ./build/web-service/ ./cmd/web-service/main.go

# run manually #
$ go run \\cmd\\web-service\\main.go 
```

## ❓ QnA

No QnA right now

## 👀 Known Issues

No issues right now

## 💪 Todo
1. Move translation responsibility into service layer (moved from rest.createResponse function)
2. Refactor storage.saveFile using os Write (remove fasthttp dependency)
3. Enhancment storage.SaveFile using date foldering
4. Complete Unit Test
5. Refactor Mock using `Testify` (if necessary)
6. End to end Test (Test usecase)
7. Database implementation `mongodb`
8. MySQL Database Support

## 🤩 Nice to Have
1. `AWS S3` Support
2. `Built In` Support (using other instance of go-seidon)
3. Concurrent/pararell processing when uploading multiple files
4. Custom file validation rules (e/g: based on provider)
5. Caching layer
6. Allowing file authorization in the future.
7. Custom file slug configuration.
8. `Alibabab` Support
9. Storage dashboard monitoring
10. Multiple provider for each file support (e.g: for backup purpose)
11. Dependency injection

## 💖 Contributions

Please always follow the development guidance described above to keep the `code quality` great and also to `decrease unecessary bug`. 

![Test coverage][coverage-image]

Do run the `unit test` after changing code or before push/updating the code.

> *Leave better than you found it*

[goseidon-logo]: asset/image/go-seidon.png?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
[architecture-image]: asset/image/system-architecture.jpg?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
