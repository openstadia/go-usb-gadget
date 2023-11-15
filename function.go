package gadget

const FunctionsDir = "functions"

type Function interface {
	Path() string
	Name() string
}
