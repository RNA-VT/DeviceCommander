// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewDevice struct {
	Mac         *string `json:"MAC"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        string  `json:"Host"`
	Port        int     `json:"Port"`
	Active      *bool   `json:"Active"`
}

type UpdateDevice struct {
	ID          string  `json:"ID"`
	Mac         *string `json:"MAC"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        *string `json:"Host"`
	Port        *int    `json:"Port"`
	Active      *bool   `json:"Active"`
}