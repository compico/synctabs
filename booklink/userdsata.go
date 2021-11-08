package booklink

import "errors"

type UserData struct {
	Bookmarks []Bookmark `json:"bookmarks"`
	User      *User      `json:"user"`
	Setting   *Setting   `json:"setting"`
}

func InitUserData() *UserData {
	ud := new(UserData)
	ud.Bookmarks = make([]Bookmark, 0)
	return ud
}

func (ud *UserData) AddBookmarks(itsfolder bool, icon string, name string, title string, description string, images []string, url string) {
	bm := Bookmark{
		ItsFolder: itsfolder,
	}
	bm.Sitelinks = make([]SiteData, 0)
	bm.Sitelinks = append(bm.Sitelinks, SiteData{
		Icon:        icon,
		Name:        name,
		Title:       title,
		Description: description,
		Images:      images,
		Url:         url,
	})
	ud.Bookmarks = append(ud.Bookmarks, bm)
}

func (ud *UserData) AddFolder(name string, icon string, marks []SiteData) error {
	bm := Bookmark{
		Name:      name,
		Icon:      icon,
		ItsFolder: true,
	}
	bm.Sitelinks = make([]SiteData, 0)
	emptyFolder := false
	if marks == nil {
		emptyFolder = true
	}
	if emptyFolder {
		if len(marks) <= 0 {
			return errors.New("Send nil marks for empty folder!")
		}
		bm.Sitelinks = append(bm.Sitelinks, marks...)
	}
	ud.Bookmarks = append(ud.Bookmarks, bm)
	return nil
}
