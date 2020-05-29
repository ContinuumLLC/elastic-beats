// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package system

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded gzipped contents of module/system.
func AssetSystem() string {
	return "eJy0WU1v27gW3ftXXHTTBHAVxG2CwosHpC8PL8G002CSAt3ZlHgtcUyRGpJKov76AakPSzZlW44qIEBEk+ec+0VeSR9gjcUcdKENphMAwwzHObx7dAPvJgAUdaRYZpgUc/jPBADgKUGNQBSCSRBWDDnVEKNARQxSCAs3XmJCKmnOMZgAKORINM4hREMmUC2cTyYAH0CQFOeAzyiM4zBFhnOIlcwzd19Ptv/Xs6ViMRNuqF6wxuJFKlqNebTb67tbB3LldDrOAJ4SpiEiAkIEAivGETJiEjjDIA5gefFM1AWXsf0LLpfn0wZNKgdjJdWQlemRTDMpUBgwCTGg8yzjDKmbQokhNbZAw5lYL8+Dti9yjepoV6AwzBQLRod74/4WcsH+yZEXwKgFWhVMxE6l1QBSAIFEahPAvQHrJZlmuY000UDg8e7mw+zqGhKik41TSkfYVXB/Oy2B7D9E0PLG6g46NhhUKROEDzfhqVpZ01qCji8zJSPU+mh3tmzZnt4r4o7oBHWTVa8Y5YaEHG1qobVDu5IhPJaKmSR1VNo5xC54JjxHN6VBdB7EV0ARSYoUKItRm2qms29b/8aCkJM1zsLF7Op6g+fx6JY5X77e/PG/WdgE1GPOpIfp4+dPpzB9/PxpKNPV5ewUpqvL2bFMOiGz2SBzHu9uZrOjLdEJGeiux7ubAZ6y+IvhFrg1wziGpVfJcXxuOY4TPLUY6quBKeU4huXT1eXshIhcXc4uhsXE8QyOiuM5Pi6vr8n1IFN+/rzea0RjgIzWeHwH8JuOvVLF4eOuMqbBbI67EgCYkBSnwGVEONw/1P9lUpkpKEylQTdsz4Dq1v7W9YjrJQKSU+b3i8e+7pHQOtakNs2g72jb4y97LS3AEiIpDGGi7vl4aTcTK6lSYtcFrVXbXV99bWtsNT6ZYSl2iEulXIq4M1wSzoHmyvF2fmQiy82iniKIkBojKajuzJK5aU8j+pYU3hmZwohp55TLzu97/GWvH84aYKItIfCYHUppegynxOAQzi9SGrBYPp4qeqjYL6QeslBKjkQM4Xu0ub6q0sAWScPhE2CF/ZICA3vrEbC9kRwh4M9W813Dt3vQKbhO+8vj015BcrXSaAKN0THZd0DT00aHRbUZsCf6VuV4/rir0HxMzBf0Ezng/tZHQVSUMIORydWIBnVgq2en18/Xi+tP5z4RKfFF8QTubzf/BUKpQq3RGzuWeYi2Bg9w3D/sp5DaQ7G9cx9gWUrd2rtb2zWQUObGFYvM7EO8PQerc6e73+7s2e1theJOAu/z+kGffH9sQKd2eyGiqKKujUITJeeBV0nGibG2jaqkBq0URCiM1FPIw1yYfAovTFD5onsUje4X97KhVPKNRHbkZw/1iqSMF6OSl5AVvUKaEDMFiiEjYgorhRhqesgjz6j09oH9Vl0Vpp9wjUogH4/vyVMs73VFsyulyU0SrUmMb2rCKoy9tUwEMKEN4RwpSOW6y2ekNf/bGrTtrvuQM/e6cl8fXqsd3IjXV9OQV0hQbia2365HPFnTW7YnmvjQIvfx+IrhjVR7rKriPSZbBdnXCoxJ1e4BfHycRSjGta6C9J78ZY2N0rbXdBVmb/+u2a+jHo6OIrNgXpI8TYkqTgAsF/owc8XHDMuPv77u7q/NS/M2xZDN1QIc7JLsJF2+F99tk47fT39XgwDwo/uGfcdLbBvx7WzdJ4ENVzwu1/9tLHvJKFNjG/ZeQyJTtNAYGdlN7fabN+QjthcAD0rGiqRgJKhcADHAZcx6uhubkItWro7q8eolj/tq037JA98FfGUif52CSZi2J7Qtjhgjqcts78mInceWWqEM/8bIDBO4dHAHmqGiJNWbj1pMQ0aUsY3DWYiFrL7C5GXEM8XsLlau2mph/ZUM+6v5UBSOigQ0+b9b2rC35Db0TBiMcbtKBtL3lV9GtPYY1/e0eji2NeD+8DZRq2bDmZCmaiCrEWY08tXgSFrlvyuSNzuyLWwAD1JrFvL2F0FY6oRQ+bJo/NGDedYx2nXGtjBF+VbbYbhP2+fTjW8XlGkScqTLaQ/qUsgNs+Uoi50SEaOSuXb9uCikQPcBncsYmDh3bXYfYqSKzLRBXxIU3ZC52FjtF2iiCzdMQSOmugfUyDpL7OMPCsfhnnlKxJ3ot7pGos0iSqxB/aWz086V11HBfnKf/IvOHlMb+kK0EwCVgGDybwAAAP//v4QD6w=="
}