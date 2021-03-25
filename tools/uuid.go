package tools

import (
	"encoding/binary"
	guuid "github.com/google/uuid"
    "github.com/satori/go.uuid"
)

//type UID struct {}

func GenerateUUID32() uint32{
	id := guuid.New().ID()
	return id
}

func GenerateUUID64() int64{
	uid:= uuid.NewV4()
	res := binary.BigEndian.Uint64(uid[:8])
	return int64(res)

}