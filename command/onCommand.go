package command


type OffCommand struct {
	device Device
}


func(oc *OffCommand) execute() {
	oc.device.off()
}
