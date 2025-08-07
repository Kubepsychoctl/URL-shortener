package link

import (
	"app/url-shorter/internal/stat"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	URL   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := Link{
		URL: url,
	}
	link.GenerateHash()
	return &link
}

func (link *Link) GenerateHash() {
	link.Hash = RandomHash(64)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomHash(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
