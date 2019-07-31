package webaudio

// Code generated by cdproto-gen. DO NOT EDIT.

// EventContextCreated notifies that a new BaseAudioContext has been created.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextCreated
type EventContextCreated struct {
	Context *BaseAudioContext `json:"context"`
}

// EventContextWillBeDestroyed notifies that an existing BaseAudioContext
// will be destroyed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextWillBeDestroyed
type EventContextWillBeDestroyed struct {
	ContextID ContextID `json:"contextId"`
}

// EventContextChanged notifies that existing BaseAudioContext has changed
// some properties (id stays the same)..
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#event-contextChanged
type EventContextChanged struct {
	Context *BaseAudioContext `json:"context"`
}
