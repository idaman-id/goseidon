# Service Storage
Service for managing files on the application.

## 🚀 Motivations
1. Provide single point of entry to upload file.
2. Provide single point of entry to accessing the uploaded file.
3. Allowing multiple files upload at once.
4. Allowing multiple storage `provider` such as: `Built in` or `AWS S3`
5. Support multiple `language` interface, currently supported: `id`, `en`
6. Allowing file authorization in the future (next release)

## 📋 Documentation Index
- [**API Documentation**](doc/API.md)
- [**Business Documentation**](doc/BUSINESS.md)
- [**Database Documentation**](doc/DATABASE.md)
- [**Testing Documentation**](doc/TESTING.md)
- [**Relevant Document**](doc/DOCUMENT.md)

## 🖖 Dependencies
### Inbound Dependencies
1. Client with a given `application` configuration (app_key, app_secret)
2. Public resource through `/file/:id` endpoint

### Outbound Dependencies
1. Built in file service
2. AWS S3

## ✔️ Running the App

### Setup
- Copy `.env.example` into `.env`
- Setup `.env` configuration

### Development
```
$ go run main.go
```

### Deployment

Adjust deployment according to production pipeline, e.g: using `docker`, etc.
But in general we can simply run with the following command:

```
$ go run main.go
```

## ❓ QnA

No QnA right now

## 👀 Known Issues

No issues right now

## 💪 Todo

Nothing todo right now

## 🤩 Nice to Have

Nothing to have right now

## 💖 Contributions

Please always follow the development guidance described above to keep the `code quality` great and also to `decrease unecessary bug`. 

Do run the `unit test` after changing code or before push/updating the code.

> *Leave better than you found it*

[coverage-image]: asset/image/test-coverage.png?raw=true
