package component

//Igniter - Thesr are hot... when we want them to be
type Igniter struct {
	BaseComponent `yaml:",inline"`
}

func (i *Igniter) String() string {
	return ""
}

func (i *Igniter) State() string {
	return ""
}

func (i *Igniter) Enable(bool) {

}

func (i *Igniter) Disable() {

}
