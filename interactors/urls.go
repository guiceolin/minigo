package interactors

import (
	"log"
	"strconv"

	"github.com/guiceolin/minigo/models"
	"github.com/guiceolin/minigo/repositories"
	"github.com/xor-gate/bjf"
)

type UrlInteractor struct {
	Repo repository.UrlRepository
}

func (i UrlInteractor) CreateUrl(url string) string {
	var id = i.Repo.Store(models.Url{Original: url, Count: 0})

	var idToEncode = strconv.FormatUint(uint64(id), 10)
	return bjf.Encode(idToEncode)
}

func (i UrlInteractor) FindShortURL(shortURL string) models.Url {
	id := bjf.Decode(shortURL)
	url, err := i.Repo.GetById(uint64(id))
	if err != nil {
		log.Fatal(err)
	}
	return *url
}

func (i UrlInteractor) IncrementAccess(url models.Url) {
	url.Count++
	i.Repo.Update(url)
}
