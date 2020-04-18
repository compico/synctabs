package main

type navbarData struct {
	Url string
}

func getNavbarData(url string) *navbarData {
	data := new(navbarData)

	data.Url = url
	return data
}
