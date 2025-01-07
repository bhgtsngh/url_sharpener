package storage

import (
    "net/url"
    "strings"
    "sync"
    "sort"
)

var (
    urlMap    = make(map[string]string)
    shortMap  = make(map[string]string)
    domainMap = make(map[string]int)
    mu        sync.Mutex
)

func SaveURL(original, short string) {
    mu.Lock()
    defer mu.Unlock()

    urlMap[original] = short
    shortMap[short] = original

    domain := extractDomain(original)
    domainMap[domain]++
}

func GetShortURL(original string) (string, bool) {
    mu.Lock()
    defer mu.Unlock()
    short, exists := urlMap[original]
    return short, exists
}

func GetOriginalURL(short string) (string, bool) {
    mu.Lock()
    defer mu.Unlock()
    original, exists := shortMap[short]
    return original, exists
}

func GetTopDomains(n int) map[string]int {
    mu.Lock()
    defer mu.Unlock()

    
    domainList := make([]struct {
        domain string
        count  int
    }, 0, len(domainMap))

    for domain, count := range domainMap {
        domainList = append(domainList, struct {
            domain string
            count  int
        }{domain, count})
    }

    
    sort.Slice(domainList, func(i, j int) bool {
        return domainList[i].count > domainList[j].count
    })

    
    topDomains := make(map[string]int)
    for i := 0; i < n && i < len(domainList); i++ {
        topDomains[domainList[i].domain] = domainList[i].count
    }

    return topDomains
}

func extractDomain(urlStr string) string {
    
    parsedURL, err := url.Parse(urlStr)
    if err != nil {
        return "" 
    }

    
    domain := parsedURL.Hostname()

    
    if strings.HasPrefix(domain, "www.") {
        domain = domain[4:]
    }

    return domain
}
