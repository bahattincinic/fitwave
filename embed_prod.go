//go:build prod

package fitwave

import "embed"

//go:embed ui/dist
var UI embed.FS
