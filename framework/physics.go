package framework

//Collision is a struct holding information about a collision between two square sprites
type Collision struct {
	Time                     float64
	PositionX                float64
	Positiony                float64
	CoefficientOfRestitution float64
}

//linearEquation is used by the next function to find collisions
type linearEquation struct {
    Gradient float64
    Intercept float64
}

//findCollisionWith finds a point of collision between two sprites
// the 's sprite' and the 'o other reference sprite' does not
func (s *Sprite) findCollisionWith(o *Sprite) (collides bool, collision *Collision) {

	// Here we take a mean average of the two coefficients of restitution
	// https://en.wikipedia.org/wiki/Coefficient_of_restitution
	// This is roughly in line with how real solid body physics works
	CoefficientOfRestitution := (s.CoefficientOfRestitution + o.CoefficientOfRestitution) / 2

    // The types of collisions we will consider fall broadly into three categories:
    // - corner collisions, of which there are four
    // - edge-on-edge collisions, of which there are also held to be four
    // - stickiness, 'initial' or 'resting' collisions
    // We do not consider rotations here.
    //
    // To find the collisions, we first form a series of eight linear equations in t (for time)
    // Each linear equation represents the position of one of the edges of the two sprites
    // 
    // sx1   sx2         ox1   ox2
    //    ---   sy1         ---   oy1
    //   |   |             |   |
    //    ---   sy2         ---   oy2
    // 
    // To discover a corner collision it is necessary 
    // - for two edge equations in the X axis to share a solution at a time 't0'
    // - while two edge equations in the Y axis also share the same solution 't0'
    //
    // We are only interested in solutions in the time interval (0,1]
    //
    // To discover an edge solution it is necessary
    // - for two equations in a given axis to have a solution at some time 't0'
    // - while at this time, 't0', there is 'overlap' of the two sprites in the other axis.
    //
    // Given that the two sprites may be of different size: There are at least 9 different
    // classes of overlap that can occur in a given axis. So to assert the existence of an
    // overlap it is more efficient to assert the falsehood of the compliment.
    //
    // The two important statements are:
    // - If the first sprite is neither fully above, nor fully below the second sprite;
    //   there must be an overlap in the Y axis.
    // - If the first sprite is neither fully to the left of, nor fully to the right of
    //   the second sprite; there must be an overlap in the X axis.

    // First we define the set of eight linear equations
    sx1 := &linearEquation{s.SpeedX, s.PositionX}
    sx2 := &linearEquation{s.SpeedX, (s.PositionX + s.Width)}
    sy1 := &linearEquation{s.SpeedY, s.PositionY}
    sy2 := &linearEquation{s.SpeedY, (s.PositionY + s.Height)}
    ox1 := &linearEquation{o.SpeedX, o.PositionX}
    ox2 := &linearEquation{o.SpeedX, (o.PositionX + o.Width)}
    oy1 := &linearEquation{o.SpeedY, o.PositionY}
    oy2 := &linearEquation{o.SpeedY, (o.PositionY + o.Height)}

    // Then we solve the relevant pairs of equations to find the time at which the faces line up
    // A collision with s to the right of o
    t_sx1_ox2 := linearEquationSolution(sx1, ox2)
    // A collision with s to the left of o
    t_sx2_ox1 := linearEquationSolution(sx2, ox1)
    // A collision with s below o
    t_sy1_oy2 := linearEquationSolution(sy1, oy2)
    // A collision with s below o
    t_sy1_oy2 := linearEquationSolution(sy1, oy2)
}

//collisionTime solves a pair of linear equations that determines the time of intersection
func linearEquationSolution(s, o *linearEquation) (intersectionTime float64) {

	// if the two gradients/speeds are equal they will not solve/collide
	// we use minus one as it is not in the interesting interval
	if s.Gradient == o.Gradient {
		return -1
	}

	// s(time) = (s.Gradient * time) + s.Intercept
	// o(time) = (o.Gradient * time) + o.Intercept
	// In a solution/collision s == o
	// =>
	// (s.Gradient * time) + s.Intercept == (o.Gradient * time) + o.Intercept
	// =>
	// time == (o.Intercept - s.Intercept) / (s.Gradient - o.Gradient)
	return (o.Intercept - s.Intercept) / (s.Gradient - o.Gradient)
}
