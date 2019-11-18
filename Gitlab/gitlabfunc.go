package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	gitlaboauth "golang.org/x/oauth2/gitlab"
)

var (
	gitlabConfig = &oauth2.Config{

		ClientID:     "b87b60524a7f39f8ba0178b68f3ebe281fb9cc9ac8ecc7a3fa1b52ed5729dfa7",
		ClientSecret: "cad4f0669032eb9dac1847bb7f84b9745c5f6911ddd919a07e4a2cec1d84c41c",

		Scopes:   []string{"api"},
		Endpoint: gitlaboauth.Endpoint,
	}
	randoms = "true"
)

func handlegitlab(w http.ResponseWriter, r *http.Request) {
	url := gitlabConfig.AuthCodeURL(randoms)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// redirect from Bitbucket with a token
func handlegitlabCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != randoms {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", randoms, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := gitlabConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://gitlab.example.com/api/v4/user?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	if err != nil {
		fmt.Printf("could not parse response '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprint(w, "response:  ", content)
}
