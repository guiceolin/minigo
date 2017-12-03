package interactors

import (
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
