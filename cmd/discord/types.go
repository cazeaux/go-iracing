package discord

import "time"

type DiscordWebhook struct {
	Content    string       `json:"content,omitempty"`
	Embeds     []Embeds     `json:"embeds,omitempty"`
	Components []Components `json:"components,omitempty"`
}
type Image struct {
	URL string `json:"url,omitempty"`
}
type Thumbnail struct {
	URL string `json:"url,omitempty"`
}
type Fields struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}
type Footer struct {
	Text string `json:"text,omitempty"`
}
type Embeds struct {
	Title       string     `json:"title,omitempty"`
	Type        string     `json:"type,omitempty"`
	Color       int        `json:"color,omitempty"`
	Description string     `json:"description,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Fields      []Fields   `json:"fields,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
}
type Component struct {
	Type  int    `json:"type,omitempty"`
	Style int    `json:"style,omitempty"`
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}
type Components struct {
	Type       int         `json:"type,omitempty"`
	Components []Component `json:"components,omitempty"`
}
