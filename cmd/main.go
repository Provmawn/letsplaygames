package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/* none of this works
type Response struct {
	GamesCount int    `json:"games_count"`
	Games      []Game `json:"games"`
}

type Game struct {
	AppId                  int `json:"appid"`
	PlaytimeForever        int `json:"playtime_forever"`
	PlaytimeWindowsForever int `json:"playtime_windows_forever"`
	PlaytimeMacForever     int `json:"playtime_mac_forever"`
	PlaytimeLinuxForever   int `json:"playtime_linux_forever"`
	RTimeLastPlayed        int `json:"rtime_last_played"`
}

*/

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func GetOwnedGames(token string, steamid string) map[string]interface{} {
	// do api request to steam
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + token + "&steamid=" + steamid + "&format=json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer r.Body.Close()

	// print out the body
	bodyBytes, err := io.ReadAll(r.Body)
	fmt.Println(len(string(bodyBytes)))

	// decode the response
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)
	return result
}

func GetAppList() map[string]interface{} {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get("http://api.steampowered.com/ISteamApps/GetAppList/v2/")
	// do api request to steam
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer r.Body.Close()

	// print out the body
	bodyBytes, err := io.ReadAll(r.Body)
	fmt.Println(len(string(bodyBytes)))

	// decode the response
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)
	return result
}

func main() {
	token := flag.String("token", "", "steam api token")
	steamid := flag.String("steamid", "", "steam ID")
	flag.Parse()
	games := GetOwnedGames(*token, *steamid)
	fmt.Println(PrettyPrint(games))
	fmt.Println(GetAppList())

}
