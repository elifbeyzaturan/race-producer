package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/elifbeyzaturan/race-producer/internal/kafka"
	"github.com/elifbeyzaturan/race-producer/internal/randomizer"
)

const (
	kafkaBroker  = "localhost:9092"
	tickInterval = 4 * time.Second
	raceInterval = 10 * time.Minute
)

func main() {
	producer := kafka.NewProducer(kafkaBroker)
	defer producer.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		runRace(producer, r)
		log.Println("Next race starts in 10 minutes...")
		time.Sleep(raceInterval)
	}
}

func runRace(producer *kafka.Producer, r *rand.Rand) {
	ctx := context.Background()
	race := randomizer.NewRace(r)
	log.Printf("Race started: %s | Discipline: %s", race.ID, race.Discipline.Name)

	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	for range ticker.C {
		randomizer.Tick(&race, r)

		if err := producer.Send(ctx, race); err != nil {
			log.Printf("Failed to send message: %v", err)
		}

		log.Printf("Tick %d sent | Elapsed: %dms", race.Tick, race.ElapsedMs)

		if randomizer.IsFinished(&race) {
			race.Status = "FINISHED"
			if err := producer.Send(ctx, race); err != nil {
				log.Printf("Failed to send final message: %v", err)
			}
			log.Printf("Race finished: %s | Discipline: %s", race.ID, race.Discipline.Name)
			return
		}
	}
}
