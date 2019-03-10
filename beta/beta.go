package beta

import (
    "github.com/maxb/go-dep-version-test/alpha"
)

func Beta() string {
    return "beta" + alpha.Alpha()
}
