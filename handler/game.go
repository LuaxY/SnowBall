package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"

    "github.com/dustin/go-humanize"
)

type Game struct {
}

type Player struct {
    Name     string `json:"name"`
    ServerId uint16 `json:"serverId"`
    Level    uint16 `json:"level"`
    Kamas    uint64 `json:"kamas"`
}

func (h *Game) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(""))

    var player Player

    decoder := json.NewDecoder(r.Body)

    if err := decoder.Decode(&player); err != nil {
        return
    }

    if err := player.Store(); err != nil {
        panic(err)
    }
}

func (p *Player) Store() error {
    file, err := os.OpenFile("public/players.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

    if err != nil {
        return err
    }

    defer file.Close()

    _, err = file.WriteString(p.String())

    return err
}

func (p *Player) String() string {
    var serverName string

    switch p.ServerId {
    case 36:
        serverName = "Agride"
    case 201:
        serverName = "Echo"
    case 202:
        serverName = "Crocabulia"
    case 203:
        serverName = "Rubilax"
    case 204:
        serverName = "Atcham"
    case 205:
        serverName = "Meriana"
    case 206:
        serverName = "Pandore"
    case 207:
        serverName = "Ush"
    case 208:
        serverName = "Julith"
    case 209:
        serverName = "Nidas"
    case 210:
        serverName = "Merkator"
    case 211:
        serverName = "Furye"
    case 212:
        serverName = "Brumen"
    case 222:
        serverName = "Ilyzaelle"
    case 50:
        serverName = "Ombre"
    case 22:
        serverName = "Oto Mustam"
    case 223:
        serverName = "Temporis I"
    case 224:
        serverName = "Temporis II"
    default:
        serverName = fmt.Sprintf("Unknown (%d)", p.ServerId)
    }

    return fmt.Sprintf("%.19s | %13s | %20.20s | %4d | %13s\r\n", time.Now(), serverName, p.Name, p.Level, humanize.Comma(int64(p.Kamas)))
}
