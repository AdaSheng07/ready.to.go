package main

type elevatorDeploymentService struct {
	d *Deploy
}

func (svc *elevatorDeploymentService) DeployElevator(elevator *Elevator) string {
	svc.d.Operation(elevator)
	return svc.d.deployStrategy
}
