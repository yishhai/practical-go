package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

/*
This is replaced with an anonymous struct inside the getGithubInfo function

	type Response struct {
		Name string `json:"name,omitempty"`
		// this struct field name does not match a particular JSON field from our response
		// but Go will still add it because of our tag
		NumOfRepos int `json:"public_repos,omitempty"`
	}
*/
func main() {

	n, nr, err := getGithubInfo("Jesserc")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Printf("Github details:\n Username: %v, Repository count: %v\n", n, nr)

}

func getGithubInfo(name string) (string, int, error) {
	// PathEscape escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.
	n := url.PathEscape(name)
	url := fmt.Sprintf("https://api.github.com/users/%s", n)
	fmt.Printf("url: %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		// log.Fatalf("error %s", err)
		/*
			The code above is equal to
			fmt.Printf("Error %v", err)
			os.Exit(1)
		*/
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		// log.Fatalf("error: %s", resp.Status)
		return "", 0, fmt.Errorf("%v - %s", url, resp.Status)
	}

	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// This code will print the response body to the console in a non-formatted form
	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	log.Fatalf("error: %s", err)
	// }

	/*
		RELATING JSON TYPES TO GO TYPES

		JSON <-> GO
		string <-> string
		null <-> nil
		number <-> float64, however Go still have float32, int8, int16, int32, int64, int, uint8,...and so on
		array <-> []any or []interfaces{}(old version). this is because arrays in Json can be mixed with different types so we use a generic type in terms of Go
		object <-> map[string]any or struct

		Go has the Time type but Json does not

	*/

	// anonymous struct
	// this is possible cause the struct isn't used outside this function
	var r struct {
		Name string `json:"name,omitempty"`
		// this struct field name does not match a particular JSON field from our response
		// but Go will still add it because of our tag
		NumOfRepos int `json:"public_repos,omitempty"`
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		// log.Fatalf("error: %s", err)
		return "", 0, err

	}
	defer resp.Body.Close()

	return r.Name, r.NumOfRepos, nil
}
