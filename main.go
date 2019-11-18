package main

import (
	"Integration_engine/bitbucket"
	"Integration_engine/github"
	"fmt"
	_ "io/ioutil"
	"net/http"
)

// create a Link for GITHUB and BITBUCKET in html page
const htmlIndex = `<html><body>
Logged in with <a href="/logingit">GitHub</a>
<br><BR><BR><Br><BR>Logged in with <a href="/loginbitbucket">Bitbucket</a>
<br><BR><BR><Br><BR>Logged in with <a href="/logingitlab">GitLab</a>
</body></html>`

// main function of the web page
func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/logingit", github.HandleGitHubLogin)

	http.HandleFunc("/loginbitbucket", bitbucket.HandleBitbucketLogin)
	//	http.HandleFunc("/logingitlab", handlegitlab)
	http.HandleFunc("/welcomegit", github.HandleGitHubCallback)
	http.HandleFunc("/welcomebitbucket", bitbucket.HandlebitbucketCallback)
	//	http.HandleFunc("/welcomegitlab", handlegitlabCallback)
	fmt.Print("Started running on http://127.0.0.1:9000\n")
	fmt.Println(http.ListenAndServe(":9000", nil))

}

//Main API or url through which we can access the github and other account
func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}
