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

    soccerball := &framework.Sprite{
        assets["soccerball"],
        20,
        35,
        9,
        16,
        10,
        8,
    }

    framework.WindowAnimationFrame(moveFootball(renderer, soccerball))
    framework.ClockedAnimationFrame(100, updateSoccerBall(soccerball))

    return true
}

func moveFootball(renderer *framework.Renderer, soccerball *framework.Sprite) framework.AnimationCallback {

    return func(...interface{}) interface{} {

        renderer.AddSprite(soccerball)
        renderer.Render()
        framework.WindowAnimationFrame(moveFootball(renderer, soccerball))
        return true
    }
}

func updateSoccerBall(soccerball *framework.Sprite) framework.AnimationCallback {

    return func(...interface{}) interface{} {

        // make the ball bounce off the walls
        if soccerball.PositionX <= 0 {
            soccerball.SetXSpeed(-soccerball.SpeedX)
        }
        if soccerball.PositionX >= (100 - soccerball.Width) {
            soccerball.SetXSpeed(-soccerball.SpeedX)
        }
        if soccerball.PositionY <= 0 {
            soccerball.SetYSpeed(-soccerball.SpeedY)
        }
        if soccerball.PositionY >= (100 - soccerball.Height) {
            soccerball.SetYSpeed(-soccerball.SpeedY)
        }
        soccerball.Update()

        framework.ClockedAnimationFrame(100, updateSoccerBall(soccerball))
        return true
    }
}
