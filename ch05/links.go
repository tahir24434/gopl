package links
import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"
)

// Extracts makes an HTTP GET request to the specified URL, pareses the response
// as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
    }
    
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }
    
    var links []string
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key == "href" {
                    link, err := resp.Request.URL.Parse(a.Val)
                    if err != nil {
                        continue
                    }
                    links = append(links, link.String())
                }
            }
        }
    }
    forEachNode(doc, visitNode, nil)
    return links, err
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }
    
    for c := n.FirstChild; c != nil; c = n.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        post(n)
    }
}
