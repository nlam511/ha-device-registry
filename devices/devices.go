package devices

type Device struct {
	Name string
	Type string
	Ip   string
}

type Devices struct {
	devices []Device
}

func New() *Devices {
	return &Devices{
		devices: []Device{},
	}
}

func (d *Devices) GetAll() []Device {
	return d.devices
}

func (d *Devices) Add(device Device) {
	d.devices = append(d.devices, device)
}
