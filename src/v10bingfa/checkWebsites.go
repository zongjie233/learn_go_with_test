package v10bingfa

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		go func() { // 启用多线程
			results[url] = wc(url)
		}()
	}
	return results
}
