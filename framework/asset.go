package framework

import "syscall/js"

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
		assets[name] = NewAsset(uri)
	}

	return
}

//PercentageofAssetsLoaded loops over the assets and checks that they have finished
// loading or failing to load
func FractionOfAssetsLoaded(assets map[string]*Asset) float64 {

	count := 0

	for _, asset := range assets {
		if asset.Image.Get("complete").Bool() {
			count += 1
		}
	}

	return float64(count) / float64(len(assets))
}
