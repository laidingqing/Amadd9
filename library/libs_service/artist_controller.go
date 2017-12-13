package libs_service

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	. "github.com/laidingqing/amadd9/common/entities"
	. "github.com/laidingqing/amadd9/common/services"
	"gopkg.in/mgo.v2/bson"
)

type ArtistController struct{}

var artistUri = "/artists"

type ArtistResponse struct {
	Artist ArtistRecord `json:"artist,omitempty"`
}

type ArtistList struct {
	List []ArtistRecord `json:"ea:artists"`
}

type ArtistIndexResponse struct {
	TotalRows       int        `json:"totalRows"`
	PageNum         int        `json:"offset"`
	ArtistIndexList ArtistList `json:"_embedded"`
}

func (ac ArtistController) AddRoutes(ws *restful.WebService) {
	ws.Route(ws.POST(artistUri + "").To(ac.create).
		Doc("create new artist in library.").
		Operation("create").
		Writes(ArtistResponse{}))
}

func (ac ArtistController) create(request *restful.Request, response *restful.Response) {
	// curUser := GetCurrentUser(request, response)
	// if curUser == nil {
	// 	Unauthenticated(request, response)
	// 	return
	// }

	newArtist := new(ArtistRecord)
	err := request.ReadEntity(newArtist)
	if err != nil {
		WriteBadRequestError(response)
		return
	}

	rev, err := new(ArtistManager).Create(newArtist, nil)
	if err != nil {
		WriteError(err, response)
		return
	}
	response.AddHeader("ETag", rev)
	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(&ArtistRecord{
		ID: bson.ObjectIdHex(rev),
	})
}
