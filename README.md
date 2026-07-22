# race-producer

A Go service that simulates Olympic swimming races and publishes live race updates to Kafka every 4 seconds.

## Overview

- 8 countries, 4 swimmers each (32 swimmers total)
- 15 disciplines (individual + relay)
- A new race starts every 10 minutes
- Race updates are published to Kafka topic `race-updates` every 4 seconds

## Simulation Logic

Each tick represents 4 seconds of real wall-clock time. At the start of every tick, each swimmer is assigned a random speed (m/s) drawn from an Olympic-level range for their stroke style. Distance covered in that tick is calculated as `speed x 4`.

Positions are ranked by total distance covered. Finished swimmers are ranked by finish time.

For relay races, mid-tick leg transitions are supported: when a swimmer completes their leg within a tick, the next swimmer immediately starts with the remaining time in that tick. This means different teams' leg handoffs occur at different moments, even within the same tick.

## Disciplines

| Name | Style | Distance | Type |
|------|-------|----------|------|
| Freestyle 50m | freestyle | 50m | Individual |
| Freestyle 100m | freestyle | 100m | Individual |
| Freestyle 200m | freestyle | 200m | Individual |
| Freestyle 400m | freestyle | 400m | Individual |
| Backstroke 100m | backstroke | 100m | Individual |
| Backstroke 200m | backstroke | 200m | Individual |
| Breaststroke 100m | breaststroke | 100m | Individual |
| Breaststroke 200m | breaststroke | 200m | Individual |
| Butterfly 100m | butterfly | 100m | Individual |
| Butterfly 200m | butterfly | 200m | Individual |
| Individual Medley 200m | medley | 200m | Individual |
| Individual Medley 400m | medley | 400m | Individual |
| 4x100m Freestyle Relay | freestyle | 4x100m | Relay |
| 4x100m Medley Relay | medley | 4x100m | Relay |
| 4x200m Freestyle Relay | freestyle | 4x200m | Relay |

Medley relay leg order follows Olympic rules: backstroke → breaststroke → butterfly → freestyle.

## Speed Ranges

| Style | Min (m/s) | Max (m/s) |
|-------|-----------|-----------|
| Freestyle | 1.78 | 2.27 |
| Backstroke | 1.56 | 1.92 |
| Breaststroke | 1.39 | 1.72 |
| Butterfly | 1.61 | 2.00 |
| Medley | 1.52 | 1.92 |

## Kafka Message Format

Published to topic `race-updates` every 4 seconds.

```json
{
  "raceId": "race-1721480000000",
  "status": "IN_PROGRESS",
  "discipline": {
    "name": "Freestyle 200m",
    "style": "freestyle",
    "distance": 200,
    "isRelay": false,
    "legs": 1,
    "legDist": 200
  },
  "tick": 3,
  "elapsedMs": 12000,
  "lanes": [
    {
      "number": 1,
      "country": "Turkey",
      "swimmer": { "id": 1, "name": "Ali Yilmaz", "country": "Turkey" },
      "position": 2,
      "finished": false,
      "finishTimeMs": 0,
      "distanceCovered": 89.4,
      "totalDistance": 200
    }
  ]
}
```

When the race ends, a final message is sent with `"status": "FINISHED"`.

## How to Run

**Prerequisites:** Go 1.21+, Kafka running on `localhost:9092`

```bash
go run cmd/main.go
```
