package service

import (
	"github.com/anhoder/netease-music/util"
)

type ArtistDescService struct {
	ID string `json:"id" form:"id"`
}

func (service *ArtistDescService) ArtistDesc() (float64, []byte) {

	options := &util.Options{
		Crypto:  "weapi",
	}
	data := make(map[string]string)
	data["id"] = service.ID
	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/weapi/artist/introduction`, data, options)

	return code, reBody
}
