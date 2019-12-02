package main

import (
	"Integration_engine/github"
	"fmt"
	"html/template"
	"integration-engine/Integration_engine/Gitlab"
	"integration-engine/Integration_engine/bitbucket"
	_ "io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Html_code/*.html"))
}

func main() {
	//	c := gin.Default()

	http.HandleFunc("/", handleMain)

	http.HandleFunc("/logingit", github.HandleGitHubLogin)      // github
	http.HandleFunc("/welcomegit", github.HandleGitHubCallback) // github
	http.HandleFunc("/gitrepofetch", github.FetchRepositry)     // github
	http.HandleFunc("/gitcreate", github.Create)                //github
	http.HandleFunc("/gitdelete", github.Delete)

	http.HandleFunc("/loginbitbucket", bitbucket.HandleBitbucketLogin)      //bitbucket
	http.HandleFunc("/welcomebitbucket", bitbucket.HandlebitbucketCallback) //bitbucket
	http.HandleFunc("/bitbucketrepofetch", bitbucket.Fetchrepositry)        //bitbucket
	http.HandleFunc("/Createbitbucket", bitbucket.Create)                   //bitbucket
	http.HandleFunc("/Deletebitbucket", bitbucket.Delete)

	http.HandleFunc("/logingitlab", Gitlab.Handlegitlab)       //Standard Gitlab
	http.HandleFunc("/gitlabcallback", Gitlab.GitlabCallback)  //Standard Gitlab
	http.HandleFunc("/gitlabrepofetch", Gitlab.Fetchrepositry) //Standard GitLab
	http.HandleFunc("/gitlabcreate", Gitlab.Create)            //Standard Gitlab
	http.HandleFunc("/gitlabdelete", Gitlab.Delete)            //Standard GitLab

	fmt.Print("Started running on http:localhost:9000\n") //local host
	fmt.Println(http.ListenAndServe(":9000", nil))        //local host

}

func handleMain(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Home.html", nil)
}
