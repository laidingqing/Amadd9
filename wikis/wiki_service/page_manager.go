package wiki_service

import "github.com/laidingqing/wikifeat/wikis/wiki_service/wikit"

type PageManager struct{}

//Gets a list of pages for a given wiki
func (pm *PageManager) Index(wiki string, curUser *CurrentUserInfo) (wikit.PageIndex, error) {
	auth := curUser.Auth
	theWiki := wikit.SelectWiki(Connection, wikiDbString(wiki), auth)
	return theWiki.GetPageIndex()
}
