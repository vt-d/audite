package routes

// unfinished

import (
	"runtime"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/vt-d/audite/models"
)

func InitWebsocket(app *fiber.App) {
	app.Use("/v4/websocket", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/v4/websocket", websocket.New(websocketHandle))
}

func handleStats(c *websocket.Conn) {
	for {
		var rtm runtime.MemStats
		runtime.ReadMemStats(&rtm)

		c.WriteJSON(models.Stats{
			Players:        0,
			PlayingPlayers: 0,
			Uptime:         0,
			MemoryStats: models.MemoryStats{
				Allocated:  rtm.TotalAlloc,
				Used:       rtm.HeapInuse,
				Reservable: rtm.Sys - rtm.HeapInuse,
				Free:       rtm.Frees,
			},
		})

		time.Sleep(time.Second * 30)
	}
}

func websocketHandle(c *websocket.Conn) {
	var err error

	if err = c.WriteJSON(&models.Ready{
		Op:        "ready",
		SessionID: "...",
		Resumed:   false,
	}); err != nil {
		log.Error("Error sending websocket ready:", err)
		return
	}

	go handleStats(c)

	for {
		if _, _, err = c.ReadMessage(); err != nil {
			break
		}
	}

	log.Info("Closed ws conn")
}
