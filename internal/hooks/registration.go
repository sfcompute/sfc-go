package hooks

// initHooks is invoked by Hooks.New (in hooks.go) and registers all custom
// hooks for this SDK. This file is user-owned and preserved across Speakeasy
// regenerations.
func initHooks(h *Hooks) {
	clientHeader := &ClientHeaderHook{}
	h.registerSDKInitHook(clientHeader)
	h.registerBeforeRequestHook(clientHeader)
}
