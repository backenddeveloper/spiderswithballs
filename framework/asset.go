package framework

import "syscall/js"
import "fmt"

var assetList map[string]string = map[string]string{
    "soccerball": "assets/soccerball.svg",
}

//Assest is a wrapper around a javascript Image
type Asset struct {
    Image js.Value
}

//NewAsset takes a string uri argument and creates a new Asset
// Under the hood we create a Javascript image and let the browser handle its loading
func NewAsset(uri string) *Asset {

    image := js.Global().Get("Image").New()
    image.Set("src", uri)

    return &Asset{
       image,
    }
}

//LoadAssets will take a list of URIs and return a list of assets
func LoadAssets(uris map[string]string) (assets map[string]*Asset) {

    assets = make(map[string]*Asset)

    for name, uri := range uris {
        fmt.Println(name)
        fmt.Println(uri)
        assets[name] = NewAsset(uri)
    }

    return
}
