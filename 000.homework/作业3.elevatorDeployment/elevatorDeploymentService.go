package main

import "time"

type elevatorDeploymentService struct {
	d *Deploy
}

func (svc *elevatorDeploymentService) DeployElevator(elevator *Elevator) (int, time.Duration) {
	elevator.GetOrderOfDockedFloors(elevator)
	return svc.d.Operation(elevator)
}
