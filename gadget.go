package gadget

import (
	"fmt"
	o "github.com/openstadia/go-usb-gadget/option"
	"log"
	"os"
	"path/filepath"
)

const BasePath = "/sys/kernel/config/usb_gadget"
const StringsDir = "strings"
const LangUsEng = 0x0409

type Gadget struct {
	path string
	name string
	udc  string

	configs   []*Config
	functions []Function
}

type GadgetAttrs struct {
	BcdUSB          o.Option[uint16]
	BDeviceClass    o.Option[uint8]
	BDeviceSubClass o.Option[uint8]
	BDeviceProtocol o.Option[uint8]
	BMaxPacketSize0 o.Option[uint8]
	IdVendor        o.Option[uint16]
	IdProduct       o.Option[uint16]
	BcdDevice       o.Option[uint16]
}

type GadgetStrs struct {
	SerialNumber string
	Manufacturer string
	Product      string
}

func CreateGadget(name string) *Gadget {
	path := filepath.Join(BasePath, name)

	gadget := &Gadget{
		path: BasePath,
		name: name,
	}

	err := os.Mkdir(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return gadget
}

func (g *Gadget) Enable(udc string) {
	WriteString(g.path, g.name, "UDC", udc)
	g.udc = udc
}

func (g *Gadget) Disable() {
	WriteString(g.path, g.name, "UDC", "\n")
	g.udc = ""
}

func (g *Gadget) SetAttrs(attrs *GadgetAttrs) {
	if attrs.BcdUSB.IsSome() {
		g.writeHex16("bcdUSB", attrs.BcdUSB.Value())
	}

	if attrs.BDeviceClass.IsSome() {
		g.writeHex8("bDeviceClass", attrs.BDeviceClass.Value())
	}

	if attrs.BDeviceSubClass.IsSome() {
		g.writeHex8("bDeviceSubClass", attrs.BDeviceSubClass.Value())
	}

	if attrs.BDeviceProtocol.IsSome() {
		g.writeHex8("bDeviceProtocol", attrs.BDeviceProtocol.Value())
	}

	if attrs.BMaxPacketSize0.IsSome() {
		g.writeHex8("bMaxPacketSize0", attrs.BMaxPacketSize0.Value())
	}

	if attrs.IdVendor.IsSome() {
		g.writeHex16("idVendor", attrs.IdVendor.Value())
	}

	if attrs.IdProduct.IsSome() {
		g.writeHex16("idProduct", attrs.IdProduct.Value())
	}

	if attrs.BcdDevice.IsSome() {
		g.writeHex16("bcdDevice", attrs.BcdDevice.Value())
	}
}

func (g *Gadget) SetStrs(strs *GadgetStrs, lang int) {
	langStr := fmt.Sprintf("0x%x", lang)
	path := filepath.Join(g.path, g.name, StringsDir, langStr)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	WriteString(path, "", "serialnumber", strs.SerialNumber)
	WriteString(path, "", "manufacturer", strs.Manufacturer)
	WriteString(path, "", "product", strs.Product)
}

func (g *Gadget) writeHex16(file string, value uint16) {
	WriteHex16(g.path, g.name, file, value)
}

func (g *Gadget) writeHex8(file string, value uint8) {
	WriteHex8(g.path, g.name, file, value)
}

func (g *Gadget) writeBuf(file string, buf []byte) {
	WriteBuf(g.path, g.name, file, buf)
}

func (g *Gadget) Path() string {
	return g.path
}

func (g *Gadget) Name() string {
	return g.name
}
