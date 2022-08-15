package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

/*
	This is the client code.
	We can't change this code.
*/
type CodePanel struct {
}

/*
	Method expects KeyHolder interface  to obtain text code.
	We can't change this code.
*/
func (cp *CodePanel) EnterTextKey(key KeyHolder) {
	k := key.GetTextKey()
	fmt.Printf("CodePanel: Typing %s into the pannel.\n", k)
}

/*
	Interface used in the client code.
	We can't change this code.
*/
type KeyHolder interface {
	GetTextKey() string
}

/*
	Struct with interface supported initially by the client code.
	We can't change this code.
*/
type TextCode struct {
	TextCode string
}

// We can't change this code.
func (k *TextCode) GetTextKey() string {
	return k.TextCode
}

/*
	Struct with incompatible interface.
	We can't change this code.
*/
type VoiceKey struct {
	VoiceCode string
}

/*
	Adapter for VoiceKey struct.
	Embedding is used to implement VoiceKey interface.
*/
type VoiceKeyAdapter struct {
	VoiceKey
}

/*
	Implementation of required interface
*/
func (k *VoiceKeyAdapter) GetTextKey() string {
	fmt.Println("VoiceKeyAdapter: Converting speech into text...")
	return k.VoiceCode
}

func main() {
	client := &CodePanel{}

	gk := &TextCode{"123"}

	/*
		  Output:
		- CodePanel: Typing 123 into the pannel.
	*/
	client.EnterTextKey(gk)

	vk := &VoiceKey{"Hello!"}

	/*
		Can't execute this
		client.EnterTextKey(vk)
	*/

	va := &VoiceKeyAdapter{*vk}

	/*
		  Output:
		- VoiceKeyAdapter: Converting speech into text...
		- CodePanel: Typing Hello! into the pannel.
	*/
	client.EnterTextKey(va)
}
