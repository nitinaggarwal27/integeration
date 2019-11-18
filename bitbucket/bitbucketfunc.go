package bitbucket

import (
	"fmt"
	_ "io/ioutil"
	"net/http"

	_ "github.com/ktrysmt/go-bitbucket"
	"golang.org/x/oauth2"
	bitbucketoauth "golang.org/x/oauth2/bitbucket"
)

var (
	bitbucketConfig = &oauth2.Config{

		ClientID:     "JHfvHfj8vHSxWrPPQ9",
		ClientSecret: "UpjKmSbtApTRJTaryhf3xPkcxjjMJNKZ",
		Scopes:       []string{"account", "repository"},
		Endpoint:     bitbucketoauth.Endpoint,
	}
	randomSate = "true"
)

func HandleBitbucketLogin(w http.ResponseWriter, r *http.Request) {
	url := bitbucketConfig.AuthCodeURL(randomSate)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// redirect from Bitbucket with a token
func HandlebitbucketCallback(w http.ResponseWriter, r *http.Request) {

	token, err := bitbucketConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	state := r.FormValue("state")
	if state != randomSate {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", randomSate, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://bitbucket.org/site/oauth2/access_token" + token.AccessToken)
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprint(w, "response:  ", resp)
	code := r.FormValue("code")
	fmt.Fprint(w, "\n\n\n\n\n", token)
	fmt.Println(code)
	/*
		defer resp.Body.Close()

		   	ak, err := ioutil.ReadAll(resp.Body)
		   	//http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		   	if err != nil {
		   		fmt.Printf("could not parse response '%s'\n", err)
		   		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		   		return
		   	}

		   	fmt.Fprint(w, "response:  ", token)
		   	//tokenToJSON(token)

		   }

		   /*
		   func tokenToJSON(token *oauth2.Token) {
		   	d, _ := json.Marshal(token)
		   	fmt.Println("json token", d)
		   }*/
	//AccessToken := token.AccessToken

	// user token
	acces_token := token.AccessToken
	fmt.Println(acces_token)
	//http.Get("https://bitbucket.org/site/oauth2/access_token" + token.AccessToken)
	//ak, aks := http.Get("http://bitbucket.org/!api/2.0/user?access_token=" + acces_token)

}
