package wfap

import "log"

const ApInterfaceName = "uap0"

type ApInterface struct {
}

func (c *ApInterface) AddInterface() (string, error) {
	return executeCmd("iw", "phy", "phy0", "interface", "add", ApInterfaceName, "type", "__ap")
}

func (c *ApInterface) RemoveInterface() (string, error) {
	return executeCmd("iw", "dev", ApInterfaceName, "del")
}

func (c *ApInterface) UpInterface() (string, error) {
	return executeCmd("ifconfig", ApInterfaceName, "up")
}

func (c *ApInterface) StartInterface() {
	log.Println("start interface")
	// c.RemoveInterface()
	// c.AddInterface()
	// c.UpInterface()
}
