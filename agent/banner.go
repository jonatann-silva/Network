package agent

import (
	"fmt"

	"jonatann-silva/network/version"
)

// Banner is a banner to be displayed when the DSN
// agent is started
var Banner = fmt.Sprintf(`
====|===================>
___  ____ ____ ____ ____ 
|  \ |__/ |__| | __ |  | 
|__/ |  \ |  | |__] |__| 
		   
                  {{ .AnsiColor.Cyan }}%s{{ .AnsiColor.Default }}
<===================|====

`, version.GetVersion().VersionNumber())
