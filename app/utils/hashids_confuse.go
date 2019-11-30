package utils

import (
	"blog/app/config"
	"github.com/speps/go-hashids"
)

//混淆
func HashidsEncode(src_data int) string {
	hd := hashids.NewData()
	hd.Salt = config.HASHIDS_SALT
	h, _ := hashids.NewWithData(hd)
	des_data, _ := h.Encode([]int{src_data})
	return des_data
}

//还原
func HashidsDecode(des_data string) int {
	hd := hashids.NewData()
	hd.Salt = config.HASHIDS_SALT
	h, _ := hashids.NewWithData(hd)
	src_data, _ := h.DecodeWithError(des_data)
	return src_data[0]
}
