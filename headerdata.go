package main

type Headerdata struct {
	Title  string
	Define string
}

func getheaderdata(title string) *Headerdata {
	data := new(Headerdata)

	data.Title = title
	return data
}
