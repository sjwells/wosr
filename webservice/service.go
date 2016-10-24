package main

type authorService interface {
	getAuthorByUUID(uuid string) (author, bool)
}

type simpleAuthorService struct {
	authors map[string]author
}

func newHardCodedAuthorService() simpleAuthorService {
	return simpleAuthorService{authors: map[string]author{"123": author{
		Name:          "Martin Wolf",
		Email:         "martin.wolf@ft.com",
		ImageURL:      "https://next-geebee.ft.com/image/v1/images/raw/fthead:martin-wolf?source=next",
		Biography:     "Martin Wolf is chief economics commentator at the Financial Times, London.",
		TwitterHandle: "@martinwolf_",
		TmeIdentifier: "Q0ItMDAwMDkwMA==-QXV0aG9ycw==",
	}, "456": author{
		Name:          "Lucy Kellaway",
		Email:         "lucy.kellaway@ft.com",
		ImageURL:      "https://next-geebee.ft.com/image/v1/images/raw/fthead:lucy-kellaway?source=next",
		Biography:     "Lucy Kellaway is an Associate Editor and management columnist of the FT. For the past 15 years her weekly Monday column has poked fun at management fads and jargon and celebrated the ups and downs of office life.",
		TmeIdentifier: "Q0ItMDAwMDkyNg==-QXV0aG9ycw==",
	}}}
}

func (as simpleAuthorService) getAuthorByUUID(uuid string) (author, bool) {

	requestedAuthor, found := as.authors[uuid]
	return requestedAuthor, found
}
