// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID       int    `json:"id"`
	Comment  string `json:"comment"`
	Likes    int    `json:"likes"`
	User     *User  `json:"user"`
	PostDate string `json:"postDate"`
}

type Episode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewComment struct {
	EpisodeID int    `json:"episodeID"`
	Comment   string `json:"comment"`
	Username  string `json:"username"`
}

type Program struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type QueryEpisodesInput struct {
	ProgramID int `json:"programID"`
	From      int `json:"from"`
	To        int `json:"to"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}