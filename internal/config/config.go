package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type (
	StrConfig struct {
		Menu struct {
			Button struct {
				Back         string `yaml:"back"`
				Organization string `yaml:"organization"`
				Address      string `yaml:"address"`
				Name         string `yaml:"name"`
				View         string `yaml:"view"`
				Overview     string `yaml:"overview"`
				List         string `yaml:"list"`
				Compare      string `yaml:"compare"`
				Region       string `yaml:"region"`
				Menu         string `yaml:"menu"`
				Next         string `yaml:"next"`
				NextRow      string `yaml:"nextRow"`
				Load         string `yaml:"load"`
				Code         string `yaml:"code"`
				Intermediate string `yaml:"intermediate"`
			}
		}
		Response struct {
			MsgToUser struct {
				EnterAddress string `yaml:"enterAddress"`
				EnterName    string `yaml:"enterName"`
				EnterRegion  string `yaml:"enterRegion"`
				None         string `yaml:"none"`
				Processing   string `yaml:"processing"`
				Incorrect    string `yaml:"incorrect"`
			} `yaml:"msgToUser"`
			View struct {
				Format string `yaml:"format"`
			}
			Overview struct {
				WithMonth    string `yaml:"withMonth"`
				WithoutMonth string `yaml:"withoutMonth"`
			}
		}
		Status struct {
			ProcessingStatus struct {
				Main         string `yaml:"main"`
				Region       string `yaml:"region"`
				Organization string `yaml:"organization"`
				Address      string `yaml:"address"`
				Name         string `yaml:"name"`
				View         string `yaml:"view"`
				Overview     string `yaml:"overview"`
				Code         string `yaml:"code"`
				Intermediate string `yaml:"intermediate"`
			} `yaml:"processingStatus"`

			StatusSearch struct {
				None string `yaml:"none"`
			} `yaml:"searchStatus"`
		} `yaml:"status"`
	}

	Server struct {
		Port string `yaml:"port"`
	}

	TranscriptRespSource struct {
		Type  map[string]string `yaml:"type"`
		State map[string]string `yaml:"state"`
	}

	TGConfig struct {
		TokenTG string `env:"TG_TOKEN"`
		AppURL  string `env:"APP_URL"`
	}

	GPConfig struct {
		ClientSecret string `env:"GP_CLIENT_SECRET"`
		ClientId     string `env:"GP_CLIENT_ID"`
		UrlAuth      string `env:"GP_URL_AUTH"`
		UrlRequest   string `env:"GP_URL_REQUEST"`
		GrantType    string `env:"GP_GRANT_TYPE"`
	}

	HDLConfig struct {
		SearchUrl string `env:"SEARCH_URL"`
	}

	DBConfig struct {
		MongoDB struct {
			Host     string `env:"MONGO_HOST"`
			Port     string `env:"MONGO_PORT"`
			Username string `env:"MONGO_USERNAME"`
			Password string `env:"MONGO_PASSWORD"`
			DBName   string `env:"MONGO_DBNAME"`
			GPData   string `env:"MONGO_COLLECTION_GP"`
			Access   string `env:"MONGO_COLLECTION_ACCESS"`
		}
	}
)

func Init() (*TGConfig, *DBConfig, *GPConfig, *HDLConfig, error) {
	var (
		cfgWeb TGConfig
		cfgDB  DBConfig
		cfgGP  GPConfig
		cfgHDL HDLConfig
	)

	re := regexp.MustCompile(`^(.*` + "CaduceusTour" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		err.Error()
	}

	readEnv(&cfgWeb)
	readEnv(&cfgDB)
	readEnv(&cfgGP)
	readEnv(&cfgHDL)

	return &cfgWeb, &cfgDB, &cfgGP, &cfgHDL, nil
}

func readEnv(cfg interface{}) {
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func InitStrCfg(filename string) (StrConfig, TranscriptRespSource, Server) {
	var (
		cfgStr        StrConfig
		transcription TranscriptRespSource
		server        Server
	)

	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	unmarshalYaml(ymlFile, &cfgStr)
	unmarshalYaml(ymlFile, &transcription)
	unmarshalYaml(ymlFile, &server)

	return cfgStr, transcription, server
}

func unmarshalYaml(ymlFile []byte, out interface{}) {
	err := yaml.Unmarshal(ymlFile, out)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
