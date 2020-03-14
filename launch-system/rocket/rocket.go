package rocket

import "time"

type state interface {
	getState() string
	actuateControlSystems()
	adjustPosition()
}

type sensor struct {
	x    int
	y    int
	fuel int
}

var rocketState state = groundIdle{}

var rocketSensor = sensor{0, 0, 100}

const (
	GroundIdleState       = "GROUND_IDLE"
	PoweredFlightState    = "POWERED_FLIGHT"
	BallisticDescentState = "BALLISTIC_DESCENT"
	LandingSafeState      = "LANDING_SAFE"
)

type groundIdle struct{}
type poweredFlight struct{}
type ballisticDescent struct{}
type landingSafe struct{}

var groundControlInput chan bool

func LaunchRocket(input chan bool, output chan string) {

	groundControlInput = input

	for {
		actuateControlSystems()
		adjustPosition()

		time.Sleep(100 * time.Millisecond)
		output <- rocketState.getState()

		_, ok := rocketState.(landingSafe)
		if ok {
			return
		}
	}
}

func actuateControlSystems() {
	rocketState.actuateControlSystems()
}

func adjustPosition() {
	rocketState.adjustPosition()
}

func (groundIdle) getState() string {
	return GroundIdleState
}

func (groundIdle) actuateControlSystems() {
	m := <-groundControlInput
	if m {
		rocketState = poweredFlight{}
	}
}

func (groundIdle) adjustPosition() {}

func (poweredFlight) getState() string {
	return PoweredFlightState
}

func (poweredFlight) actuateControlSystems() {
	if rocketSensor.fuel <= 0 {
		rocketState = ballisticDescent{}
	}
}

func (poweredFlight) adjustPosition() {
	rocketSensor.fuel -= 10
	rocketSensor.y += 10
}

func (ballisticDescent) getState() string {
	return BallisticDescentState
}

func (ballisticDescent) actuateControlSystems() {
	if rocketSensor.y <= 0 {
		rocketState = landingSafe{}
	}
}

func (ballisticDescent) adjustPosition() {
	rocketSensor.y -= 10
}

func (landingSafe) getState() string {
	return LandingSafeState
}

func (landingSafe) actuateControlSystems() {}

func (landingSafe) adjustPosition() {}
