package framework

//Sprite is an object to be rendered onto the screen
type Sprite struct {
	Asset     *Asset
	PositionX float64
	PositionY float64
	Width     float64
	Height    float64
    SpeedX    float64
    SpeedY    float64
    CoefficientOfRestitution float64
    GravitationalConstant float64
}

//SetXPosition sets the Sprite's X position
func (s *Sprite) SetXPosition (x float64) {
    s.PositionX = x
}

//SetYPosition sets the Sprite's Y position
func (s *Sprite) SetYPosition (y float64) {
    s.PositionY = y
}

//SetPosition sets the Sprite's position
func (s *Sprite) SetPosition (x,y float64) {
    s.SetXPosition(x)
    s.SetYPosition(y)
}

//SetWidth sets the sprite's width
func (s *Sprite) SetWidth (width float64) {
    s.Width = width
}

//SetHeight sets the sprite's height
func (s *Sprite) SetHeight (height float64) {
    s.Height = height
}

//SetSize sets the sprite's size
func (s *Sprite) SetSize (width, height float64) {
    s.SetWidth(width)
    s.SetHeight(height)
}

//SetXSpeed sets the sprite's horizontal speed
func (s *Sprite) SetXSpeed (x float64) {
    s.SpeedX = x
}

//SetYSpeed sets the sprite's vertical speed
func (s *Sprite) SetYSpeed (y float64) {
    s.SpeedY = y
}

//ChangeGravity sets the gravitational co-efficient
func (p *Sprite) ChangeGravity(g float64) {
    s.GravitationalConstant = g
}

//ChangeElasticity sets the coefficient of restitution
func (p *Sprite) ChangeElasticity (c float64) {
    s.CoefficientOfRestitution = c
}

//Update performs one 'clock' of a Sprites position
// it is up to the consumer to determine how frequently these happen
// Update adjusts the particle's speed for gravity and bounce then updates the sprite
func (s *Sprite) Update() {

    // accelerate due to gravity. Note that the framework deals in positions as
    // percentages, with 0,0 being the top left corner and 100,100 being the bottom right
    // Therefore gravity is addative to vertical speed
    s.SpeedY += s.GravitationalConstant

    // if the sprite attempts to go below impact with the floor it bounces
    // we also check it is moving in a downwards direction to prevent a sprite 
    // rendered below the floor from vibrating.
    if s.PositionY >= (100 - s.Height) && s.SpeedY > 0 {

        // https://en.wikipedia.org/wiki/Coefficient_of_restitution
        s.SetYSpeed(-s.SpeedY * s.CoefficientOfRestitution)

    }
    // This updates the Sprite's position
    s.PositionX += p.SpeedX
    s.PositionY += p.SpeedY
}

//NewSprite creates a new Sprite
func NewSprite(asset *Asset, x, y, width, height, speedX, speedY, e, g float64) *Sprite {

    return &Sprite{asset, x, y, width, height, speedX, speedY, e, g,}
}
