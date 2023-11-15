package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	gadget "github.com/openstadia/go-usb-gadget"
	o "github.com/openstadia/go-usb-gadget/option"
	"log"
	"time"
)

func main() {
	g := gadget.CreateGadget("my_hid")

	g.SetAttrs(&gadget.GadgetAttrs{
		BcdUSB:          o.Some[uint16](0x0200),
		BDeviceClass:    o.None[uint8](),
		BDeviceSubClass: o.None[uint8](),
		BDeviceProtocol: o.None[uint8](),
		BMaxPacketSize0: o.None[uint8](),
		IdVendor:        o.Some[uint16](0x1d6b),
		IdProduct:       o.Some[uint16](0x0104),
		BcdDevice:       o.Some[uint16](0x0100),
	})

	g.SetStrs(&gadget.GadgetStrs{
		SerialNumber: "fedcba9876543210",
		Manufacturer: "Tobias Girstmair",
		Product:      "iSticktoit.net USB Device",
	}, gadget.LangUsEng)

	c := gadget.CreateConfig(g, "c", 1)

	c.SetAttrs(&gadget.ConfigAttrs{
		BmAttributes: o.None[uint8](),
		BMaxPower:    o.Some[uint8](250),
	})

	c.SetStrs(&gadget.ConfigStrs{
		Configuration: "Config 1: ECM network",
	}, gadget.LangUsEng)

	hidFunction := gadget.CreateHidFunction(g, "usb0")
	hidFunction.SetAttrs(&gadget.HidFunctionAttrs{
		Subclass:     0,
		Protocol:     0,
		ReportLength: 6,
		ReportDesc:   ReportDesc,
	})

	b := gadget.CreateBinding(c, hidFunction, hidFunction.Name())
	fmt.Println(b)

	udcs := gadget.GetUdcs()
	if len(udcs) < 1 {
		return
	}
	udc := udcs[0]

	fmt.Println(udc)

	g.Enable(udc)

	time.Sleep(5 * time.Second)

	rw, _ := hidFunction.GetReadWriter()

	m := NewMouse()

	for i := 0; i < 10; i++ {
		m.Move(3000*i, 3000*i)
		m.Update(rw.Writer)
		time.Sleep(time.Second)
	}

	g.Disable()
}

func Move(x uint16, y uint16) []byte {
	b := make([]byte, 6)
	binary.LittleEndian.PutUint16(b[1:], x)
	binary.LittleEndian.PutUint16(b[3:], y)
	return b
}

type Button uint8

const (
	Left   Button = 0
	Right  Button = 1
	Center Button = 2
)

type mouse struct {
	buttons uint8
	x       uint16
	y       uint16
	wheel   int8
}

func NewMouse() *mouse {
	return &mouse{
		buttons: 0,
		x:       0,
		y:       0,
		wheel:   0,
	}
}

func (m *mouse) Move(x int, y int) {
	m.x = uint16(x)
	m.y = uint16(y)
}

func (m *mouse) Scroll(x int, y int) {
	m.wheel = int8(y)
}

func (m *mouse) MouseDown(button Button) {
	m.buttons |= 1 << button
}
func (m *mouse) MouseUp(button Button) {
	m.buttons &= ^(1 << button)
}

func (m *mouse) Update(w *bufio.Writer) {
	report := make([]byte, 6)
	report[0] = m.buttons
	binary.LittleEndian.PutUint16(report[1:], m.x)
	binary.LittleEndian.PutUint16(report[3:], m.y)
	report[5] = byte(m.wheel)

	_, err := w.Write(report)
	if err != nil {
		log.Fatal(err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
