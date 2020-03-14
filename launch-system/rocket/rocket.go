package rocket

import (
	"fmt"
	"time"
)

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
	groundIdleState       = "GROUND_IDLE"
	poweredFlightState    = "POWERED_FLIGHT"
	ballisticDescentState = "BALLISTIC_DESCENT"
	landingSafeState      = "LANDING_SAFE"
)

type groundIdle struct{}
type poweredFlight struct{}
type ballisticDescent struct{}
type landingSafe struct{}

func LaunchRocket() {

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(rocketState.getState())

		actuateControlSystems()
		adjustPosition()

		_, ok := rocketState.(landingSafe)
		if ok {
			fmt.Println(rocketState.getState())
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
	return groundIdleState
}

func (groundIdle) actuateControlSystems() {
	rocketState = poweredFlight{}
}

func (groundIdle) adjustPosition() {}

func (poweredFlight) getState() string {
	return poweredFlightState
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
	return ballisticDescentState
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
	return landingSafeState
}

func (landingSafe) actuateControlSystems() {}

func (landingSafe) adjustPosition() {}
