package cli

import out "codehub/src/output"

const version = "1.0.0"

func Version() {
	out.Info("CodeHub version " + version)
}
