package pkg

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func GuidGeneratorInit() {
	st := sonyflake.Settings{}

	flake := sonyflake.NewSonyflake(st)

	sf = flake

}

func InstanceGuidGenerator() *sonyflake.Sonyflake {
	return sf
}

func NewGuid() (int64, error) {
	//fmt.Println("start new guid ", sf)
	id, err := sf.NextID()
	result := int64(id)
	return result, err

}
