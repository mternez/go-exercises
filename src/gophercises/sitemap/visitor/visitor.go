package visitor

import (
	"maps"
	"net/http"
	"sitemap/parser"
	"slices"
	"sync"

	"golang.org/x/net/html"
)

type SyncLinksMap struct {
	mu    sync.Mutex
	links map[string]*parser.Link
}

func (linksMap *SyncLinksMap) Put(link *parser.Link) {
	linksMap.mu.Lock()
	linksMap.links[link.Href] = link
	linksMap.mu.Unlock()
}

func (linksMap *SyncLinksMap) Get(key string) (*parser.Link, bool) {
	linksMap.mu.Lock()
	defer linksMap.mu.Unlock()
	val, ok := linksMap.links[key]
	return val, ok
}

func (linksMap *SyncLinksMap) Range(rangeFunction func(key string, value *parser.Link)) {
	linksMap.mu.Lock()
	defer linksMap.mu.Unlock()
	for key, value := range linksMap.links {
		rangeFunction(key, value)
	}
}

func (linksMap *SyncLinksMap) GetValues() []*parser.Link {
	return slices.Collect(maps.Values(linksMap.links))
}

func initSyncLinksMap() *SyncLinksMap {
	linksMap := SyncLinksMap{}
	linksMap.links = make(map[string]*parser.Link)
	return &linksMap
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
