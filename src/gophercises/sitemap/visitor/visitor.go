package visitor

import (
	"net/http"
	"sitemap/parser"
	"sync"

	"golang.org/x/net/html"
)

type SyncLinksMap struct {
	mu    sync.Mutex
	links map[string]*parser.Link
}

func (cache *SyncLinksMap) Put(link *parser.Link) {
	cache.mu.Lock()
	cache.links[link.Href] = link
	cache.mu.Unlock()
}

func (cache *SyncLinksMap) Get(key string) (*parser.Link, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	val, ok := cache.links[key]
	return val, ok
}

func (cache *SyncLinksMap) Range(rangeFunction func(key string, value *parser.Link)) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for key, value := range cache.links {
		rangeFunction(key, value)
	}
}

func initSyncLinksMap() *SyncLinksMap {
	cache := SyncLinksMap{}
	cache.links = make(map[string]*parser.Link)
	return &cache
}

func Visit(host string) *SyncLinksMap {

	linksMap := initSyncLinksMap()

	homePageLinks := getLinksInPage(host, "")

	var wg sync.WaitGroup

	for _, link := range homePageLinks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			linksMap.Put(link)
			visitInternalLinks(link, host, linksMap)
		}()
	}

	wg.Wait()

	return linksMap
}

func visitInternalLinks(link *parser.Link, host string, linksMap *SyncLinksMap) {

	if !link.Internal {
		return
	}

	linksMap.Put(link)

	linksInPageMap := getLinksInPage(host, link.Href)

	var wg sync.WaitGroup

	for key, value := range linksInPageMap {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, ok := linksMap.Get(key)
			if !ok && value.Internal {
				linksMap.Put(value)
				visitInternalLinks(value, host, linksMap)
			}
		}()
	}

	wg.Wait()
}

func getLinksInPage(host string, path string) map[string]*parser.Link {

	// TODO : handle case where host doesn't end with "/"
	resp, err := http.Get(host + "/" + path)

	if err != nil {
		panic("Couldn't open link")
	}

	page, err := html.Parse(resp.Body)

	if err != nil {
		panic("Couldn't read response body")
	}

	return parser.GetLinks(page, host)
}
