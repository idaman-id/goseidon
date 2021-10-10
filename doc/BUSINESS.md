# Business Documentation

## Glosariums
1. Client is all possible actor on the system who have `application credential` (key and secret).
2. `Storage provider` or later is known as `provider` is application which storing the physical files uploaded by the client.

## Environments
1. `APP_URL` = http://storage.domain.tld
2. `APP_DEFAULT_LOCALE` = id
3. `MIN_UPLOADED_FILE` = 1
4. `MAX_UPLOADED_FILE` = 5
5. `MIN_FILE_SIZE` = 1
6. `MAX_FILE_SIZE` = 134217728

## Rules
1. Client can upload minimum of `MIN_UPLOADED_FILE` in one single upload.
2. Client can upload maximum of `MAX_UPLOADED_FILE` in one single upload.
3. Client can upload minimum `MIN_FILE_SIZE` file size. Default is `1` bytes which mean the file should be a valid file (non zero file size)
4. Client can upload maximum of `MAX_FILE_SIZE`. Default is `134217728` bytes or `128MB`
5. [NOT IMPLEMENTED] Client can optionally specify the storage `provider` when uploading the files.
6. If `provider` is specified, all files should be uploaded to the specified provider.
7. [NOT IMPLEMENTED] If `provider` is not specified, all files should be uploaded through the `active` provider in the system according to the provider `priority` (higher value higher priority)
8. [NOT IMPLEMENTED] Client can't upload the file to the `inactive` provider.
9. [NOT IMPLEMENTED] Client should specifying the `app_key` and `app_secret` when uploading the files or checking the file detail.
10. [NOT IMPLEMENTED] Client can't upload the file using the `inactive` application.

## Constraints
1. System may partially failing the file upload process (some files uploaded in the batch may be success/fail).
2. System may reject all uploaded files when there is one or invalid file.
