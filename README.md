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
- [**Builtin App**](doc/builtin-app/README.md)
- ❌ [**Gateway App**](doc/gateway-app/README.md)

## 👷🏻 Architecture
![System Architecture][architecture-image]

## ❓ QnA

Q1. Is there any possibility for file size to be negative? 

A1. No, and for overflow problem system is gonna mark negative size as zero

---

Q2. Prefer one line code or positive case at the end of code?

A2. positive case at the end of code

## 👀 Known Issues

No issues right now

## 💪 Todo
1. [storage-local] date foldering (strategy)
2. [translation] implementation on service layer
3. [all-packages] unit test
4. [builtin-app] end to end test
5. [gateway-app] implementation (gin/echo)
6. [storage] `AWS S3` Support
7. [file] refactor multipart.Fileheader depedency to avoid coupling
8. [gateway-app] end to end test

## 🤩 Nice to Have
1. [repository] `mongodb` database implementation
2. [database] `mongodb` database client
3. [gateway-app] Concurrent processing when uploading multiple files
4. [storage] `Alicloud OSS` Support
5. [gateway-app] Custom file validation rules/policies (e.g: based on `provider`, or `application`)
6. [gateway-app] Caching layer
7. [gateway-app] Allowing file authorization in the future (e.g: based on `context`)
8. [gateway-app] Custom file slug configuration (for SEO purpose)
9. [gateway-app] Storage dashboard monitoring (e.g: grafana dashboard by using prometheus exporter)
10. [gateway-app] Upload to multiple provider for each file support (e.g: for backup purpose)

## 💖 Contributions

Please always follow the development guidance described above to keep the `code quality` great and also to `decrease unecessary bug`. 

![Test coverage][coverage-image]

Do run the `unit test` after changing code or before push/updating the code.

> *Leave better than you found it*

[goseidon-logo]: asset/image/go-seidon.png?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
[architecture-image]: asset/image/system-architecture.jpg?raw=true
[coverage-image]: asset/image/test-coverage.png?raw=true
