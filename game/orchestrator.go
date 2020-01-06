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

	ballWorld := framework.NewWorld()
	renderer.AddWorld(ballWorld)

	soccerball := framework.NewSprite(assets["soccerball"], 0, 0, 4.5, 8, 0.2, 0.01, 0.6, 0.01)
	football := framework.NewSprite(assets["football"], 100, 0, 4.5, 8, -0.2, 0.01, 0.6, 0.01)
	ballWorld.AddSprite(soccerball)
	ballWorld.AddSprite(football)

	go updateBalls(ballWorld)

	return true
}

func updateBalls(w *framework.World) {

	for {
		time.Sleep(5e5)
		w.Update()
	}
}
