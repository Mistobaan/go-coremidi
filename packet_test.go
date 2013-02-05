package coremidi

import "testing"

func TestSend(t *testing.T) {
	devices, _ := AllDevices()
	device := devices[0]
	client, _ := NewClient("test client")
	port, _ := NewOutputPort(client, "test port")
	entities, _ := device.Entities()
	destinations, _ := entities[0].Destinations()
	destination := destinations[0]
	packet := NewPacket(0x90, 0x30, 100)

	_, err := packet.Send(port, destination)

	if err != nil {
		t.Fatalf("failed to send MIDI")
	}
}