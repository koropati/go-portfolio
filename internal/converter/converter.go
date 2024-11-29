package converter

import (
	"github.com/google/uuid"
)

func StringToUUID(uuidString string) (dataUUID uuid.UUID, err error) {
	dataUUID, err = uuid.Parse(uuidString)
	if err != nil {
		return
	}
	return
}
