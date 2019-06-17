package framework

//Particle is an extension of Sprite that obeys the laws of gravity and restitution
type Particle struct {
    Sprite *Sprite
    CoefficientOfRestitution float64
    GravitationalConstant float64
}

//ChangeGravity sets the gravitational co-efficient
func (p *Particle) ChangeGravity(g float64) {
    p.GravitationalConstant = g
}

//ChangeElasticity sets the coefficient of restitution
func (p *Particle) ChangeElasticity (c float64) {
    p.CoefficientOfRestitution = c
}
