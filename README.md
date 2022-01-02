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
$ air # default configuration (.air.toml)
$ air -c .air.toml # custom configuration

# run manually #
$ go run \\cmd\\file-service\\main.go 

# testing #
$ ginkgo watch -r -v

```

### Deployment

Adjust deployment according to production pipeline, e.g: using `docker`.
But in general we can simply run with the following command:

```bash
# build manually #
$ go build -o ./build/file-service/ ./cmd/file-service/main.go

# run manually #
$ go run \\cmd\\file-service\\main.go 
```

## ❓ QnA

Q1. Is there any possibility for file size to be negative? 

A1. No, and for overflow problem system is gonna mark negative size as zero

## 👀 Known Issues

No issues right now

## 💪 Todo
1. Make all code work again after refactor
2. `mongodb` database implementation + test
3. `mysql` database implementation + test
4. Refactor storage.SaveFile using os Write (remove fasthttp dependency)
5. Enhancment storage.SaveFile using date foldering
6. Implement translation on service layer
7. Test json marshalling (struct tag)
8. End to end Test (Test usecase)

## 🤩 Nice to Have
1. `Built In` Support (using other `instance` of go-seidon)
2. Concurrent/pararell processing when uploading multiple files
3. `AWS S3` Support
4. `Alicloud OSS` Support
5. Custom file validation rules/policies (e.g: based on `provider`, or `application`)
6. Caching layer
7. Allowing file authorization in the future (e.g: based on `user context`)
8. Custom file slug configuration (for SEO purpose)
9. Storage dashboard monitoring (e.g: grafana dashboard by using prometheus exporter)
10. Upload to multiple provider for each file support (e.g: for backup purpose)

## 💖 Contributions

Please always follow the development guidance described above to keep the `code quality` great and also to `decrease unecessary bug`. 

![Test coverage][coverage-image]

Do run the `unit test` after changing code or before push/updating the code.

> *Leave better than you found it*

[goseidon-logo]: asset/image/go-seidon.png?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
[architecture-image]: asset/image/system-architecture.jpg?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
