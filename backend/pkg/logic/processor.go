package logic

import (
	"fmt"

	"github.com/Childebrand94/takeHomePhoneNumber/pkg/models"
)

func ProcessData(data models.Query) {
	fmt.Printf("this is the request: %+v\n", data)
}
