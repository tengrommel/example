package config

import "flag"

var DataDirectory = flag.String("data-directory", "", "Path for loading templates and migration scripts.")
