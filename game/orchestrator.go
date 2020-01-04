package game

import "../framework"
import "time"

//Start is the main entrypoint for the game
func Start(canvas *framework.Canvas) bool {

	// await for the assets to load
	assets := framework.LoadAssets(assetList)

	for {
		if framework.FractionOfAssetsLoaded(assets) == 1 {
			break
		}
		time.Sleep(1e7)
	}

	renderer := framework.NewRenderer(canvas)
	renderer.RenderForever()

	soccerball := framework.NewSprite(assets["football"], 100, 0, 4.5, 8, -0.2, 0.01, 0.6, 0.01)
	renderer.AddSprite(soccerball)

	go updateSoccerBall(soccerball)

	return true
}

func updateSoccerBall(soccerball *framework.Sprite) {

	for {
		time.Sleep(5e6)
		soccerball.Update()
	}
}
