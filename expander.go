package url_expander

import "net/http"

func ExpandUrl(url string) (string, error) {
	expandedUrl := url

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			expandedUrl = req.URL.String()
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return expandedUrl, nil
}
