package animegoorg_check

import (
	"anime-bot-schedule/models"
	"anime-bot-schedule/pkg/message"
	repositories_animes "anime-bot-schedule/repositories/animes"
	repositories_subscribe "anime-bot-schedule/repositories/subscribe"
	parsing "anime-bot-schedule/services/parser/animego.org"
	"fmt"
	"log"
)

func Check(anime models.Anime) {

	resp, err := parsing.Fetch(anime.URL)
	if err != nil {
		log.Printf("error fetching anime data: %s", err)
		return
	}

	var lastEpisod parsing.Episod

	if resp.Episods[0].Relized {
		lastEpisod = resp.Episods[0]
	} else if resp.Episods[1].Relized {
		lastEpisod = resp.Episods[1]
	} else if resp.Episods[2].Relized {
		lastEpisod = resp.Episods[2]
	}

	if lastEpisod.Number != anime.LastReleasedEpisode {

		subscribers, err := repositories_subscribe.GetByAnime(anime.ID)

		if err != nil {
			return
		}

		for _, subscriber := range subscribers {
			text := fmt.Sprintf("%s \n\nВышла новая серия на телеэкранах японии: %s (%s)", *resp.Title, lastEpisod.Number, lastEpisod.Title)

			msg := message.NewMessage{
				Text:        text,
				Photo:       *resp.Image,
				UserId:      subscriber.TelegramID,
				Link:        anime.URL,
				AnimeId:     anime.ID,
				LinkTitle:   "Перейти",
				DeletePrev:  true,
				Unsubscribe: true,
			}

			msg.Send()
		}

		repositories_animes.UpdateLastEpisod(anime.ID, lastEpisod.Number)
	}
}
