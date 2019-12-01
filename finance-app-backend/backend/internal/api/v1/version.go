package v1

import (
	"awesomeProject/finance-app-backend/backend/internal/config"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API for returning version
// When server starts, we set version and than use it if necessary

// ServerVersion represents the server version
type ServerVersion struct {
	Version string `json:"version"`
}

// Marshaled JSON
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}
}

// VersionHandler serves version information
func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version.")
	}
}
