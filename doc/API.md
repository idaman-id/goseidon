# API Documentation

## Symbol
- ✅ Function available
- ❌ Function unavailable
- ☑️ Testcase available
- ⚠️ Testcase unavailable

| Feature  | Symbol | Description         |
| ---      | ------ | ------              |
| Function |   ✔️  | Function available   |
| Function |   ❌  | Function unavailable |
| Testcase |   ☑️  | Testcase available   |
| Testcase |   ⚠️  | Testcase unavailable |

## General Response
All available API below may be returning general response according to the specific situation occured.

**Failed Response**
- HttpCode: 500
- Response Body: 
```json
{
	"message": "Terjadi kesalahan"
}
```

**NotFound Response**
- HttpCode: 404
- Response Body: 
```json
{
	"message": "Resource tidak ditemukan"
}
```

**Unauthenticated Response**
- HttpCode: 401
- Response Body: 
```json
{
	"message": "Akses tidak valid"
}
```

---

## Index
- [**Upload File ❌⚠️** ](#upload-file)
- [**File Detail ❌⚠️** ](#file-detail)
- [**File Resource ❌⚠️** ](#file-resource)

### Upload File
- Method: **POST**
- Endpoint: **/v1/file**
- Status: ❌⚠️

**Request Headers**
```json
{
	"Content-Type": "multipart/form-data"
}
```

**Request Body**
```json
{
	"files": [ // required, min: 1, max: 5
		// FileObject{}
		// FileObject{}
	],
	"provider": "provider_id" // optional, must be an active provider_id
}
```

**Success Response**
- HttpCode: 200
- Response Body:
```json
{
	"code": "200",
	"message": "Berhasil mengunggah file",
	"data": [
		{
			"status": "success",
			"file": {
				"uuid": "651fd093-03cb-4ff4-a23c-7959ce07def5",
				"name": "samplevideo-1280x720-1mb.mp4",
				"size": 1055736,
				"type": "video",
				"extension": "mp4",
				"mimetype": "video/mp4",
				"url": "http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4"
			}
		},
		{
			"status": "failed"
		}
	]
}
```

**Failed Response**
- HttpCode: 400
- Response Body: 
```json
{
	"message": "Terjadi kesalahan saat mengunggah file"
}
```

**Invalid Data Response**
- HttpCode: 422
- Response Body: 
```json
{
	"message": "Terdapat data yang tidak valid",
	"error": [
		{
			"field": "files.0",
			"message": "Jenis file tidak didukung"
		}
	]
}
```

---

### File Detail
- Method: **GET**
- Endpoint: **/v1/file/:id**
- Status: ❌⚠️
- Example: **http://storage.idaman.local/v1/file/651fd093-03cb-4ff4-a23c-7959ce07def5**

**Success Response**
- HttpCode: 200
- Response Body:
```json
{
	"data": {
		"uuid": "651fd093-03cb-4ff4-a23c-7959ce07def5",
		"name": "samplevideo-1280x720-1mb.mp4",
		"size": 1055736,
		"type": "video",
		"extension": "mp4",
		"mimetype": "video/mp4",
		"url": "http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4"
	}
}
```

**Failed Response**
- HttpCode: 404
- Response Body: 
```json
{
	"message": "File tidak ditemukan"
}
```

---

### File Resource
- Method: **GET**
- Endpoint: **/file/{:id}.{extension}**
- Status: ❌⚠️
- Example: **http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4**

**Success Response**
- HttpCode: 200
- Response Body: **FileObject**

**Failed Response**
- HttpCode: 404
- Response Body: **NotFound FileObject**

---
