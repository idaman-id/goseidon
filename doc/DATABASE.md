# Database Documentation
Database used by this application is `mongodb`

## Colection Index
- [File](#collection-file)
- [Provider](#collection-provider)
- [Application](#collection-application)

### Collection: File
- Collection Name: `file`
- Data Structure
```json
{
  "_id": {
    "type": "ObjectId",
    "required": true,
    "unique": true,
    "example": {
      "$oid": "606437e1ea400000f70055ce"
    }
  },
  "uuid": {
    "type": "String",
    "required": true,
    "unique": true,
    "example": "651fd093-03cb-4ff4-a23c-7959ce07def5"
  },
  "name": {
    "type": "String",
    "required": true,
    "example": "samplevideo-1280x720-1mb.mp4",
    "min": 3,
    "max": 2048
  },
  "size": {
    "type": "Uint32",
    "required": true,
    "example": 1055736
  },
  "type": {
    "type": "String",
    "required": true,
    "enum": [
      "audio", "video", "image", "document"
    ],
    "example": "video"
  },
  "extension": {
    "type": "String",
    "required": true,
    "example": "mp4",
    "min": 1,
    "max": 64
  },
  "mimetype": {
    "type": "String",
    "required": true,
    "example": "video/mp4",
    "min": 1,
    "max": 128
  },
  "url": {
    "type": "String",
    "required": true,
    "example": "http://storage.idaman.local/file/651fd093-03cb-4ff4-a23c-7959ce07def5.mp4",
    "min": 3,
    "max": 2048
  },
  "path": {
    "type": "String",
    "required": true,
    "example": "file/custom/directory/2021/03/31/samplevideo-1280x720-1mb-606437e13acba.mp4",
    "min": 3,
    "max": 2048
  },
  "provider_id": {
    "type": "ObjectId",
    "required": true,
    "example": {
      "$oid": "604e0620c276172ad04aa3c2"
    }
  },
  "application_id": {
    "type": "ObjectId",
    "required": true,
    "example": {
      "$oid": "604e0620c276172ad04aa3c2"
    }
  }
}
```

### Collection: Provider
- Collection Name: `provider`
- Data Structure
```json
{
  "_id": {
    "type": "ObjectId",
    "required": true,
    "unique": true,
    "example": {
      "$oid": "604e0620c276172ad04aa3c2"
    }
  },
  "type": {
    "type": "String",
    "required": true,
    "enum": [
      "built_in", "aws_s3"
    ],
    "example": "built_in"
  },
  "priority": {
    "type": "Uint8",
    "required": true,
    "description": "higher value is higher priority",
    "example": 1,
    "min": 1,
    "max": 32
  },
  "code": {
    "type": "String",
    "required": true,
    "unique": true,
    "example": "suzaku-1",
    "min": 3,
    "max": 100
  },
  "name": {
    "type": "String",
    "required": true,
    "example": "Suzaku 1",
    "min": 3,
    "max": 100
  },
  "host": {
    "type": "String",
    "required": true,
    "example": "http://file-service.idaman.local",
    "min": 3,
    "max": 2048
  },
  "status": {
    "type": "String",
    "required": true,
    "enum": [
      "active", "inactive"
    ],
    "example": "active"
  }
}
```

### Collection: Application
- Collection Name: `application`
- Data Structure
```json
{
  "_id": {
    "type": "ObjectId",
    "required": true,
    "unique": true,
    "example": {
      "$oid": "604e0620c276172ad04aa3c2"
    }
  },
  "name": {
    "type": "String",
    "required": true,
    "example": "Gateway Android",
    "min": 3,
    "max": 100
  },
  "key": {
    "type": "String",
    "required": true,
    "unique": true,
    "example": "jhkXcgVyszUGzE6uTKFX5ivwjxB83LQ8v+9U7NRz6f56lkxzD9CHSp",
    "max": 1024
  },
  "secret": {
    "type": "String",
    "required": true,
    "example": "uxTsn9xqmtRF27euKzKcfH/EvXHU9sVvXd+2suplySX",
    "max": 4096
  },
  "status": {
    "type": "String",
    "required": true,
    "enum": [
      "active", "inactive"
    ],
    "example": "active"
  }
}
```