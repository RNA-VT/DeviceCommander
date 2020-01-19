package component

//Igniter - Thesr are hot... when we want them to be
type Igniter struct {
	BaseComponent `yaml:",inline"`
}

func (i *Igniter) String() string {
	return ""
}

//State returns a string represnting the current state
func (i *Igniter) State() string {
	return ""
}

//Enable - enable this igniter and optionally initalize its gpio
func (i *Igniter) Enable(init bool) {

}

//Disable - disable this igniter and set state to 'off'
func (i *Igniter) Disable() {

}
