package github

import (
	"context"
	"fmt"
	"net/http"
	_ "strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	githubConfig = &oauth2.Config{

		ClientID:     "5a40d8065418bb33dae2",
		ClientSecret: "574ab4412490ab812b0aa299ac02b850578d8858",
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     githuboauth.Endpoint,
	}
	randomState = "true"
)

func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := githubConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != randomState {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", randomState, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := githubConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	/*resp, err := http.Get("https://github.com/login/oauth/access_token" + token.AccessToken)
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
		}*/
	// user code

	AccessToken := token.AccessToken

	// user token
	fmt.Println(token.AccessToken)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	fmt.Fprint(w, repos)

}
