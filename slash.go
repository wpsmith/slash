package slash

import (
	"net/http"
	"net/url"

	"golang.org/x/net/context"
)

// Command represents an incoming Slash Command request.
type Command struct {
	Token string

	TeamID     string
	TeamDomain string

	ChannelID   string
	ChannelName string

	UserID   string
	UserName string

	Command string
	Text    string

	ResponseURL *url.URL
}

// CommandFromValues returns a Command object from a url.Values object.
func CommandFromValues(v url.Values) (Command, error) {
	u, err := url.Parse(v.Get("response_url"))
	if err != nil {
		return Command{}, err
	}

	return Command{
		Token:       v.Get("token"),
		TeamID:      v.Get("team_id"),
		TeamDomain:  v.Get("team_domain"),
		ChannelID:   v.Get("channel_id"),
		ChannelName: v.Get("channel_name"),
		UserID:      v.Get("user_id"),
		UserName:    v.Get("user_name"),
		Command:     v.Get("command"),
		Text:        v.Get("text"),
		ResponseURL: u,
	}, nil
}

// ParseRequest parses the form an then returns the extracted Command.
func ParseRequest(r *http.Request) (Command, error) {
	err := r.ParseForm()
	if err != nil {
		return Command{}, err
	}
	return CommandFromValues(r.Form)

}

// Params returns the match groups from a regular expression match.
func Params(ctx context.Context) map[string]string {
	params, ok := ctx.Value(paramsKey).(map[string]string)
	if !ok {
		return make(map[string]string)
	}
	return params
}

func WithParams(ctx context.Context, params map[string]string) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}

// key used to store context values from within this package.
type key int

const (
	paramsKey key = 0
)
