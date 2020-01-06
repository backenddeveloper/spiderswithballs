package framework

//resolveCollisionWith finds a point of collision between two sprites
// the 's sprite' has its position adjusted, the 'o other reference sprite' does not
func (s *Sprite) resolveCollisionWith(o *Sprite) {

	// Here we take a mean average of the two coefficients of restitution
	// https://en.wikipedia.org/wiki/Coefficient_of_restitution
	// This is roughly in line with how real solid body physics works
	CoefficientOfRestitution := (s.CoefficientOfRestitution + o.CoefficientOfRestitution) / 2

	// first we find collisions in the X axes
	// there are 2 possible collisions that could occur,
	// "S on the left of O" and "S on the right of O"
	collisionLeftX := collisionTime((s.PositionX + s.Width), s.SpeedX, o.PositionX, o.SpeedX)
	collisionRightX := collisionTime(s.PositionX, s.SpeedX, (o.PositionX + o.Width), o.SpeedX)

	// then we find collisions in the Y axes
	// there are 2 possible collisions that could occur,
	// "S above O" and "S below O"
	collisionAboveY := collisionTime((s.PositionY + s.Height), s.SpeedY, o.PositionY, o.SpeedY)
	collisionBelowY := collisionTime(s.PositionY, s.SpeedY, (o.PositionY + o.Height), o.SpeedY)
}

//collisionTime solves a pair of linear equations that determines the time of intersection
func collisionTime(sPosition, sSpeed, oPosition, oSpeed float64) (intersectionTime float64) {

	// POSITION_s(time) = (s.Speed * time) + s.Position
	// POSITION_o(time) = (o.Speed * time) + o.Position
	// In a collision POSITION_s == POSITION_o
	// =>
	// (s.Speed * time) + s.Position == (o.Speed * time) + o.Position
	// =>
	// time == (o.Position - s.Position) / (s.Speed - o.Speed)
	return (oPosition - sPosition) / (sSpeed - oSpeed)
}

//earliestCollision finds the earliest collision out of 2 in the time interval (0, 1]
func earliestCollision(l, r float64) (collision float64) {

	// if both are in the interval
	if l > 0 && l < 1 && r > 0 && r < 1 {

		// if right is smaller
		if r < l {
			return r

			// else left is smaller or both are equal
		} else {
			return l
		}

		// if  just l is in the interval
	} else if l > 0 && l < 1 {
		return l

		// if just r is in the interval
	} else if r > 0 && r < 1 {
		return r

		// else neither are in the interval
	} else {
		return -1
	}

}
