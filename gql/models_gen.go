// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

type NewPage struct {
	Text   string `json:"text"`
	Name   string `json:"name"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name      string `json:"name"`
	AccountID string `json:"accountId"`
	UID       string `json:"uid"`
}

type UpdateUser struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Twitter     *string `json:"twitter"`
	Instagram   *string `json:"instagram"`
	Facebook    *string `json:"facebook"`
	Homepage    *string `json:"homepage"`
	Birthday    *string `json:"birthday"`
}
