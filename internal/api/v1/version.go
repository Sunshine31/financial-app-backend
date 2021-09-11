package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Sunshine31/financial-app/internal/config"
	"github.com/sirupsen/logrus"
)

type ServerVersion struct {
	Version string `json:"version"`
}

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

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version.")
	}
}
