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
All API available below may be returning general response according to the specific situation occured.

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
	"message": "Akses tidak diotentikasi"
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

**Request Body**
```json
{
	"files": [ // multipart/form-data
		// FileObject{}
	],
	"provider": "provider_id" // optional registered provider_id
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
			"uuid": "651fd093-03cb-4ff4-a23c-7959ce07def5",
			"name": "samplevideo-1280x720-1mb.mp4",
			"size": 1055736,
			"type": "video",
			"extension": "mp4",
			"mimetype": "video/mp4",
			"url": "http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4",
			"path": "file/custom/directory/2021/03/31/samplevideo-1280x720-1mb-606437e13acba.mp4",
		}
	]
}
```

**Failed Response**
- HttpCode: 400
- Response Body: 
```json
{
	"message": "Gagal mengunggah file"
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
		"url": "http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4",
		"path": "file/custom/directory/2021/03/31/samplevideo-1280x720-1mb-606437e13acba.mp4",
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
