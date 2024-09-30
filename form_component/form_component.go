package formcomponent

type Component string
const (
	Checkbox Component = "<input type='checkbox'/>"
	Range Component    = "<input type='range'/>"
	Textarea Component = "<textarea/>"
)

func InputComponent(cs []Component) string {
	form := "<form>"
	for _, v := range cs {
		form += string(v)
	}
	return form
}
