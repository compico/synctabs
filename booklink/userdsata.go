package booklink

type UserData struct {
	Bookmarks []SiteData `json:"bookmarks"`
	User      *User      `json:"user"`
	Setting   *Setting   `json:"setting"`
}

func (ud *UserData) AddBookmarks(icon string, name string, title string, description string, images []string, url string) {
	sd := SiteData{
		Icon:        icon,
		Name:        name,
		Title:       title,
		Description: description,
		Images:      images,
		Url:         url,
	}
	ud.Bookmarks = append(ud.Bookmarks, sd)
}