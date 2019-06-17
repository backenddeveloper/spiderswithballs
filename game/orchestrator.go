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

    soccerball := framework.NewSprite(assets["soccerball"], 20, 35, 9, 16, 10, 8, 1, 0.6)

    framework.WindowAnimationFrame(
        func(...interface{}) interface{} {
            renderer.AddSprite(soccerball)
            renderer.Render()
            framework.WindowAnimationFrame(moveFootball(renderer, soccerball))
            return true
        })
    go updateSoccerBall(soccerball)

    return true
}

func updateSoccerBall(soccerball *framework.Particle) {

    for {
        time.Sleep(1e7)
        soccerball.Update()
    }
}
