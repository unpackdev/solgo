package ir

import (
	"regexp"
	"strings"
)

// Link represents a link extracted from a comment.
type Link struct {
	Location string `json:"location"`  // The actual URL or link location.
	Social   bool   `json:"is_social"` // Indicates if the link is from a social media/network.
	Network  string `json:"network"`   // Represents the network or platform of the link.
}

// GetLocation returns the link's location.
func (l *Link) GetLocation() string {
	return l.Location
}

// IsSocial returns true if the link is from a social media/network.
func (l *Link) IsSocial() bool {
	return l.Social
}

// GetNetwork returns the network or platform of the link.
func (l *Link) GetNetwork() string {
	return l.Network
}

// processLinks processes the links in the comments of a RootSourceUnit.
func (b *Builder) processLinks(root *RootSourceUnit) {
	// Define a regex pattern to match URLs.
	urlPattern := `https?://[\w\d\-./:%#?=&@]+`
	urlRegex := regexp.MustCompile(urlPattern)

	links := []*Link{}

	// Loop through each comment and extract URLs.
	for _, comment := range root.GetAST().GetComments() {
		matches := urlRegex.FindAllString(comment.GetText(), -1)
		for _, match := range matches {

			match = strings.TrimSuffix(match, ".")

			link := &Link{
				Location: match,
			}

			// Check the type of platform/network of the URL.
			switch {
			case isX(match):
				link.Social = true
				link.Network = "twitter_x"
			case isTelegram(match):
				link.Social = true
				link.Network = "telegram"
			case isFacebook(match):
				link.Social = true
				link.Network = "facebook"
			case isGitHub(match):
				link.Social = true
				link.Network = "github"
			case isGitLab(match):
				link.Social = true
				link.Network = "gitlab"
			case isReddit(match):
				link.Social = true
				link.Network = "reddit"
			case isMedium(match):
				link.Social = true
				link.Network = "medium"
			case isLinkedIn(match):
				link.Social = true
				link.Network = "linkedin"
			case isDiscord(match):
				link.Social = true
				link.Network = "discord"
			case isBitcointalk(match):
				link.Social = true
				link.Network = "bitcointalk"
			case isYouTube(match):
				link.Social = true
				link.Network = "youtube"
			case isClubhouse(match):
				link.Social = true
				link.Network = "clubhouse"
			}

			links = append(links, link)
		}
	}

	root.Links = links
}

// Helper functions to determine the platform/network of a URL.

func isX(url string) bool {
	twitterPattern := `(?i)^https?://(www\.)?(twitter\.com|x\.com)/.*`
	twitterRegex := regexp.MustCompile(twitterPattern)
	return twitterRegex.MatchString(url)
}

func isTelegram(url string) bool {
	twitterPattern := `(?i)^https?://(www\.)?(telegram\.com|t\.me)/.*`
	twitterRegex := regexp.MustCompile(twitterPattern)
	return twitterRegex.MatchString(url)
}

func isFacebook(url string) bool {
	facebookPattern := `(?i)^https?://(www\.)?(facebook\.com|fb\.me)/.*`
	facebookRegex := regexp.MustCompile(facebookPattern)
	return facebookRegex.MatchString(url)
}

func isGitHub(url string) bool {
	pattern := `(?i)^https?://(www\.)?github\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isGitLab(url string) bool {
	pattern := `(?i)^https?://(www\.)?gitlab\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isReddit(url string) bool {
	pattern := `(?i)^https?://(www\.)?reddit\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isMedium(url string) bool {
	pattern := `(?i)^https?://(www\.)?medium\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isLinkedIn(url string) bool {
	pattern := `(?i)^https?://(www\.)?linkedin\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isDiscord(url string) bool {
	pattern := `(?i)^https?://(www\.)?discord\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isBitcointalk(url string) bool {
	pattern := `(?i)^https?://(www\.)?bitcointalk\.org/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isYouTube(url string) bool {
	pattern := `(?i)^https?://(www\.)?youtube\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}

func isClubhouse(url string) bool {
	pattern := `(?i)^https?://(www\.)?joinclubhouse\.com/.*`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}
