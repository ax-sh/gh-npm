package api

import gh "github.com/cli/go-gh/v2/pkg/api"

func GithubClientGet(path string, response interface{}) error {
	client, err := gh.DefaultRESTClient()
	if err != nil {
		return err
	}
	err = client.Get(path, &response)
	return err
}
