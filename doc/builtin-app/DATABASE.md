# Database Documentation
Database structure for builtin app

# MySQL Database
- Database Name: `goseidon_builtin`

## Table Index
- [File](#table-file)

### Table: File
- Table Name: `file`
- Data Structure
```json
{
  "id": {
    "type": "BigInt",
    "unsigned": true,
    "required": true,
    "unique": true,
    "primary_key": true,
    "example": 1
  },
  "unique_id": {
    "type": "Varchar",
    "required": true,
    "unique": true,
    "max": 250,
    "example": "651fd093-03cb-4ff4-a23c-7959ce07def5"
  },
  "original_name": {
    "type": "Varchar",
    "required": true,
    "example": "samplevideo 1280x720 1mb.mp4",
    "min": 3,
    "max": 512
  },
  "name": {
    "type": "Varchar",
    "required": true,
    "example": "samplevideo-1280x720-1mb.mp4",
    "min": 3,
    "max": 512
  },
  "size": {
    "type": "Int",
    "unsigned": true,
    "required": true,
    "example": 1055736
  },
  "extension": {
    "type": "Varchar",
    "required": true,
    "example": "mp4",
    "min": 1,
    "max": 32
  },
  "mimetype": {
    "type": "Varchar",
    "required": true,
    "example": "video/mp4",
    "min": 1,
    "max": 128
  },
  "file_location": {
    "type": "Varchar",
    "required": true,
    "description": "set empty string if saved on root dir",
    "example": "storage/file/custom-dir/2021/01/05",
    "min": 0,
    "max": 1024
  },
  "file_name": {
    "type": "Varchar",
    "required": true,
    "example": "651fd093-03cb-4ff4-a23c-7959ce07def5.mp4",
    "min": 3,
    "max": 512
  },
  "created_at": {
    "type": "Int",
    "unsigned": true,
    "required": true,
    "example": 1640858210
  },
  "updated_at": {
    "type": "Int",
    "unsigned": true,
    "required": false,
    "example": 1640858210
  },
  "deleted_at": {
    "type": "Int",
    "unsigned": true,
    "required": false,
    "example": 1640858210
  }
}
```

- Query Preview

```sql
  CREATE TABLE `goseidon`.`file`(  
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `unique_id` VARCHAR(250) NOT NULL,
    `name` VARCHAR(512) NOT NULL,
    `size` INT(10) UNSIGNED NOT NULL,
    `extension` VARCHAR(32) NOT NULL,
    `mimetype` VARCHAR(128) NOT NULL,
    `public_url` VARCHAR(2048) NOT NULL,
    `local_path` VARCHAR(4096) NOT NULL,
    `created_at` INT(10) UNSIGNED NOT NULL,
    `updated_at` INT(10) UNSIGNED,
    `deleted_at` INT(10) UNSIGNED,
    PRIMARY KEY (`id`)
  );
```
