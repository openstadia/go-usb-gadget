package main

import (
	"github.com/openstadia/go-usb-gadget"
	o "github.com/openstadia/go-usb-gadget/option"
)

func cdc() {
	g := gadget.CreateGadget("my_usb")

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

}
