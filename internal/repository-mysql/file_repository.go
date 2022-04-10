package repository_mysql

import (
	"database/sql"

	app_error "idaman.id/storage/internal/error"
	"idaman.id/storage/internal/file"
	"idaman.id/storage/internal/repository"
)

type fileRepository struct {
	db          *sql.DB
	fileService file.FileService
}

func (r *fileRepository) FindByIdentifier(identifier string) (*repository.FileModel, error) {

	uniqueId := r.fileService.RemoveFileExtension(identifier)
	sqlQuery := `
		SELECT
			id, unique_id, original_name, name,
			size, extension, mimetype, file_location, file_name,
			created_at, updated_at, deleted_at
		FROM file WHERE unique_id = ?`
	fileStmt, err := r.db.Prepare(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer fileStmt.Close()

	fileModel := FileModel{}
	err = fileStmt.QueryRow(uniqueId).Scan(
		&fileModel.Id, &fileModel.UniqueId, &fileModel.OriginalName, &fileModel.Name,
		&fileModel.Size, &fileModel.Extension, &fileModel.Mimetype,
		&fileModel.FileLocation, &fileModel.FileName,
		&fileModel.CreatedAt, &fileModel.UpdatedAt, &fileModel.DeletedAt,
	)
	if err != nil {
		msg := err.Error()
		if msg == "sql: no rows in result set" {
			err = app_error.NewNotfoundError("File")
		}
		return nil, err
	}

	file := repository.FileModel{
		Id:           fileModel.Id,
		UniqueId:     fileModel.UniqueId,
		OriginalName: fileModel.OriginalName,
		Name:         fileModel.Name,
		Extension:    fileModel.Extension,
		Size:         fileModel.Size,
		Mimetype:     fileModel.Mimetype,
		FileLocation: fileModel.FileLocation,
		FileName:     fileModel.FileName,
	}
	file.SetCreatedAtFromUnixTime(fileModel.CreatedAt)

	updatedAt, err := fileModel.UpdatedAt.Value()
	isUpdatedAtValid := fileModel.UpdatedAt.Valid && err == nil
	if isUpdatedAtValid {
		file.SetUpdatedAtFromUnixTime(updatedAt.(int64))
	}

	deletedAt, err := fileModel.DeletedAt.Value()
	isDeletedAtValid := fileModel.DeletedAt.Valid && err == nil
	if isDeletedAtValid {
		file.SetDeletedAtFromUnixTime(deletedAt.(int64))
	}

	return &file, nil
}

func (r *fileRepository) Save(p repository.SaveFileParam) error {
	_, err := r.db.Exec(
		"INSERT INTO file (unique_id, original_name, name, extension, size, mimetype, file_location, file_name, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		p.UniqueId, p.OriginalName, p.Name,
		p.Extension, p.Size, p.Mimetype, p.FileLocation, p.FileName,
		p.CreatedAt.Unix(),
	)
	return err
}

func (r *fileRepository) Delete(identifier string) error {
	_, err := r.db.Exec(
		"DELETE FROM file WHERE unique_id = ?",
		identifier,
	)
	return err
}

func NewFileRepository(db *sql.DB, fileService file.FileService) *fileRepository {
	return &fileRepository{db, fileService}
}
