package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type FileSemester struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Semester int                `bson:"semester" json:"id_semester"`
	FilePath string             `bson:"file_path" json:"file_path"`
}

type FileGeneric struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Keterangan string             `bson:"keterangan" json:"keterangan"`
	FilePath   string             `bson:"file_path" json:"file_path"`
}
