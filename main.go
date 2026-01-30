package main

import dependencies "software/src/Infrestructure/Dependencies"

func main() {
	r := dependencies.InitMaestria()
	r.Run(":8000")
}