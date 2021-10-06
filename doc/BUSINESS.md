# Business Documentation

## Glosariums
1. Client is all possible actor on the system who have `application credential` (key and secret).
2. `Storage provider` or later is known as `provider` is application which storing the physical files uploaded by the client.

## Environments
1. `APP_URL` = http://storage.domain.tld
2. `APP_DEFAULT_LOCALE` = id

## Rules
1. Client can upload minimum of `1`.
2. Client can upload maximum of `5`.
3. Client can optionally specify the storage `provider` when uploading the files.
4. If `provider` is specified, all files should be uploaded to the specified provider.
5. If `provider` is not specified, all files should be uploaded through the `active` provider in the system according to the provider `priority` (higher value higher priority)
6. Client can't upload the file to the `inactive` provider.
7. Client should specifying the `app_key` and `app_secret` when uploading the files or checking the file detail.
8. Client can't upload the file using the `inactive` application.
9. 

## Constraints
1. System may partially failing the file upload process (some files uploaded in the batch may be success/fail).
2. System may reject all uploaded files when there is one or invalid file.
3. Valid file is considered as:
- File type of: `audio`, `video`, `document`, `image`, `archieve`
- Maximum size of: 16MB for `image` type
- Maximum size of: 32MB for `audio` and `document` type
- Maximum size of: 128MB for `video` and `archieve` type
4. 