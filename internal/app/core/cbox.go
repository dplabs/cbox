package core

import (
	"log"
	"path"

	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/gofrs/uuid"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const CBOX_PATH = ".cbox"

var (
	basePath string
)

func resolveInCboxDir(content string) string {
	cboxBasePath := basePath
	if cboxBasePath == "" {
		var err error
		cboxBasePath, err = homedir.Dir()
		if err != nil {
			log.Fatalf("init: could not get HOME: %v", err)
		}
	}
	return path.Join(cboxBasePath, CBOX_PATH, content)
}

func CheckCboxDir(path string) string {
	basePath = path

	cboxPath := resolveInCboxDir("")
	pkg.CreateDirectoryIfNotExists(cboxPath)

	configFile := resolveInCboxDir("config.yml")
	pkg.CreateFileIfNotExists(configFile)

	spacesPath := resolveInCboxDir("spaces")
	if pkg.CreateDirectoryIfNotExists(spacesPath) {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("init: could not generate id: %v", err)
		}
		defaultSpace := models.Space{
			Label:       DEFAULT_SPACE_ID,
			Description: DEFAULT_SPACE_DESCRIPTION,
		}
		defaultSpace.ID = id

		cbox := LoadCbox(path)
		err = cbox.SpaceCreate(&defaultSpace)
		if err != nil {
			log.Fatalf("init: could not create space: %v", err)
		}
		PersistCbox(cbox)
	}

	return cboxPath
}

func LoadCbox(path string) *models.CBox {
	basePath = path

	cbox := models.CBox{
		Spaces: []models.Space{},
	}

	spaces := SpaceListFiles()

	for _, space := range spaces {
		err := cbox.SpaceCreate(space)
		if err != nil {
			log.Fatalf("load: could not create space: %v", err)
		}
	}
	return &cbox
}

func PersistCbox(cbox *models.CBox) {
	for _, space := range cbox.Spaces {
		SpaceStoreFile(&space)
	}
}

func InitCBox(path string) {
	cboxPath := CheckCboxDir(path)

	viper.AddConfigPath(cboxPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	defaultSettings()
}

func defaultSettings() {
	viper.SetDefault("cbox.default-space", "default")
	viper.SetDefault("cbox.environment", Env)

	if viper.IsSet("cbox.environment") {
		Env = viper.GetString("cbox.environment")
	}
}
