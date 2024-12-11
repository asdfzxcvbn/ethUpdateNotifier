package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"ethUpdateNotifier/dbutil"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/asdfzxcvbn/tgmessenger"
)

var (
	apps []string // list of appstore links

	messenger *tgmessenger.Messenger

	dbctx   context.Context
	queries *dbutil.Queries
)

// load apps
func init() {
	f, err := os.ReadFile("apps.json")
	checkErr(err)

	checkErr(json.Unmarshal(f, &apps))
}

// load tg messenger
func init() {
	var err error
	messenger, err = tgmessenger.NewMessenger(botToken, ethGroupID, ethTopicID, true)
	checkErr(err)
}

func main() {
	regex := regexp.MustCompile(`(?:id)(\d{9,10})`)

	for {
		for _, applink := range apps {
			appid := regex.FindStringSubmatch(applink)[1]
			if appid == "" {
				log.Printf("couldn't find app id for %s", applink)
				continue
			}

			log.Printf("checking %s for updates in 5 seconds..", applink)
			time.Sleep(5 * time.Second)

			latestInfo, err := getLatestVersionForID(appid)
			if err != nil {
				log.Printf("couldn't get version for %s: %v", applink, err)
				continue
			}

			currentVersion, err := queries.GetCurrentVersion(dbctx, appid)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					queries.InsertVersion(dbctx, dbutil.InsertVersionParams{
						ID:      appid,
						Version: latestInfo.Version,
					})

					log.Printf("inserted version %s for new app %s", latestInfo.Version, latestInfo.Name)
					continue
				}

				log.Printf("db couldn't get version for %s: %v", latestInfo.Name, err)
			}

			if currentVersion == latestInfo.Version {
				log.Printf("no new update for app %s", latestInfo.Name)
				continue
			}

			log.Printf("found update for %s! %s -> %s", latestInfo.Name, currentVersion, latestInfo.Version)

			err = queries.UpdateVersion(dbctx, dbutil.UpdateVersionParams{
				Version: latestInfo.Version,
				ID:      appid,
			})
			if err != nil {
				log.Printf("couldn't update version in db for app %s: %s -> %s -- %v", latestInfo.Name, currentVersion, latestInfo.Version, err)
				continue
			}

			if err := messenger.SendMessage(fmt.Sprintf(
				appUpdateTemplate,
				latestInfo.Name,
				currentVersion,
				latestInfo.Version,
				fmt.Sprintf("https://apps.apple.com/app/id%s", appid),
			)); err != nil {
				log.Printf("couldn't notify telegram about app update for %s: %s -> %s -- %v", latestInfo.Name, currentVersion, latestInfo.Version, err)
			}
		}

		log.Println("checked all apps, will check again in 5 minutes..")
		time.Sleep(5 * time.Minute)
	}
}
