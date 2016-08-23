package mousetrap

import (
	"github.com/gopherjs/gopherjs/js"
)

// Mousetrap is a wrapper for Mousetrap objects
type Mousetrap struct {
	*js.Object
}

// DefaultMousetrap is the default Mousetrap object
var DefaultMousetrap = Mousetrap{js.Global.Get("Mousetrap")}

// NewFromQuerySel calls document.querySelector with the given search string and calls New() with the result.
func NewFromQuerySel(search string) Mousetrap {
	return New(js.Global.Get("document").Call("querySelector", search))
}

// New creates a new Mousetrap object
func New(obj interface{}) Mousetrap {
	return Mousetrap{DefaultMousetrap.New(obj)}
}

// Bind binds a key or a sequence of keys to run the given function.
func (mt Mousetrap) Bind(keys string, run func()) {
	mt.Call("bind", keys, run)
}

// BindMulti binds all the keys or sequences in the given array to run the given function.
func (mt Mousetrap) BindMulti(keys []string, run func()) {
	mt.Call("bind", keys, run)
}

// BindEvent binds a key or a sequence of keys to run the given function.
func (mt Mousetrap) BindEvent(keys string, run func(), on string) {
	mt.Call("bind", keys, run, on)
}

// BindMultiEvent binds all the keys or sequences in the given array to run the given function.
func (mt Mousetrap) BindMultiEvent(keys []string, run func(), on string) {
	mt.Call("bind", keys, run, on)
}

// Unbind unbinds a single keyboard event or an array of keyboard events.
// You should pass in the key combination exactly as it was passed in originally to bind.
func (mt Mousetrap) Unbind(keys string) {
	mt.Call("unbind", keys)
}

// Trigger triggers a keyboard event within Mousetrap.
func (mt Mousetrap) Trigger(keys string) {
	mt.Call("trigger", keys)
}

// TriggerType triggers a keyboard event of the given type within Mousetrap.
func (mt Mousetrap) TriggerType(keys, typ string) {
	mt.Call("trigger", keys, typ)
}

// StopCallback is a method that is called to see if the keyboard event should fire based on if you are inside a form input field.
//
// If true is returned then it stops the callback from firing, if false it fires it.
//
// The default implementation is:
//
// 	   stopCallback: function(e, element, combo) {
//         // if the element has the class "mousetrap" then no need to stop
//         if ((' ' + element.className + ' ').indexOf(' mousetrap ') > -1) {
//             return false;
//         }
//
//         // stop for input, select, and textarea
//         return element.tagName == 'INPUT' || element.tagName == 'SELECT' || element.tagName == 'TEXTAREA' || (element.contentEditable && element.contentEditable == 'true');
//    }
func (mt Mousetrap) StopCallback(event, element, combo interface{}) {
	mt.Call("stopCallback", event, element, combo)
}

// Reset removes anything you have bound to mousetrap.
// This can be useful if you want to change contexts in your application without
// refreshing the page in your browser. You can ajax in a new page, call reset,
// and then bind the key events needed for that page.
//
// Internally mousetrap keeps an associative array of all the events to listen for
// so reset does not actually remove or add event listeners on the document.
// It just sets the array to be empty.
func (mt Mousetrap) Reset() {
	mt.Call("reset")
}

// Bind binds a key or a sequence of keys to run the given function.
func Bind(keys string, run func()) {
	DefaultMousetrap.Bind(keys, run)
}

// BindMulti binds all the keys or sequences in the given array to run the given function.
func BindMulti(keys []string, run func()) {
	DefaultMousetrap.BindMulti(keys, run)
}

// BindEvent binds a key or a sequence of keys to run the given function.
func BindEvent(keys string, run func(), on string) {
	DefaultMousetrap.BindEvent(keys, run, on)
}

// BindMultiEvent binds all the keys or sequences in the given array to run the given function.
func BindMultiEvent(keys []string, run func(), on string) {
	DefaultMousetrap.BindMultiEvent(keys, run, on)
}

// Unbind unbinds a single keyboard event or an array of keyboard events.
// You should pass in the key combination exactly as it was passed in originally to bind.
func Unbind(keys string) {
	DefaultMousetrap.Unbind(keys)
}

// Trigger triggers a keyboard event within Mousetrap.
func Trigger(keys string) {
	DefaultMousetrap.Trigger(keys)
}

// TriggerType triggers a keyboard event of the given type within Mousetrap.
func TriggerType(keys, typ string) {
	DefaultMousetrap.TriggerType(keys, typ)
}

// StopCallback is a method that is called to see if the keyboard event should fire based on if you are inside a form input field.
//
// If true is returned then it stops the callback from firing, if false it fires it.
//
// The default implementation is:
//
// 	   stopCallback: function(e, element, combo) {
//         // if the element has the class "mousetrap" then no need to stop
//         if ((' ' + element.className + ' ').indexOf(' mousetrap ') > -1) {
//             return false;
//         }
//
//         // stop for input, select, and textarea
//         return element.tagName == 'INPUT' || element.tagName == 'SELECT' || element.tagName == 'TEXTAREA' || (element.contentEditable && element.contentEditable == 'true');
//    }
func StopCallback(event, element, combo interface{}) {
	DefaultMousetrap.StopCallback(event, element, combo)
}

// Reset removes anything you have bound to mousetrap.
// This can be useful if you want to change contexts in your application without
// refreshing the page in your browser. You can ajax in a new page, call reset,
// and then bind the key events needed for that page.
//
// Internally mousetrap keeps an associative array of all the events to listen for
// so reset does not actually remove or add event listeners on the document.
// It just sets the array to be empty.
func Reset() {
	DefaultMousetrap.Reset()
}

// Pause stops Mousetrap events from firing.
//
// N.B. Requires the Pause/unpause extension!
// You can get it from https://github.com/ccampbell/mousetrap/tree/master/plugins/pause
func Pause() {
	DefaultMousetrap.Call("pause")
}

// Unpause allows Mousetrap events to fire again.
//
// N.B. Requires the Pause/unpause extension!
// You can get it from https://github.com/ccampbell/mousetrap/tree/master/plugins/pause
func Unpause() {
	DefaultMousetrap.Call("unpause")
}

// Record records keyboard shortcuts from your application.
// For example if you wanted to let users specify their own keyboard shortcuts for
// performing actions on your page you could ask them to enter a shortcut.
//
// N.B. Requires the Record extension!
// You can get it from https://github.com/ccampbell/mousetrap/tree/master/plugins/record
func Record(callback func(sequence []string)) {
	DefaultMousetrap.Call("record", callback)
}
