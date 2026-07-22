package randomizer

import "github.com/elifbeyzaturan/race-producer/internal/model"

var Countries = []model.Country{
	{
		Name: "Turkey",
		Swimmers: []model.Swimmer{
			{ID: 1, Name: "Ali Yilmaz", Country: "Turkey"},
			{ID: 2, Name: "Ayse Kaya", Country: "Turkey"},
			{ID: 3, Name: "Mehmet Demir", Country: "Turkey"},
			{ID: 4, Name: "Fatma Celik", Country: "Turkey"},
		},
	},
	{
		Name: "USA",
		Swimmers: []model.Swimmer{
			{ID: 5, Name: "Michael Johnson", Country: "USA"},
			{ID: 6, Name: "Sarah Williams", Country: "USA"},
			{ID: 7, Name: "James Brown", Country: "USA"},
			{ID: 8, Name: "Emily Davis", Country: "USA"},
		},
	},
	{
		Name: "Australia",
		Swimmers: []model.Swimmer{
			{ID: 9, Name: "Liam Smith", Country: "Australia"},
			{ID: 10, Name: "Olivia Jones", Country: "Australia"},
			{ID: 11, Name: "Noah Wilson", Country: "Australia"},
			{ID: 12, Name: "Emma Taylor", Country: "Australia"},
		},
	},
	{
		Name: "France",
		Swimmers: []model.Swimmer{
			{ID: 13, Name: "Lucas Martin", Country: "France"},
			{ID: 14, Name: "Camille Bernard", Country: "France"},
			{ID: 15, Name: "Hugo Thomas", Country: "France"},
			{ID: 16, Name: "Lea Petit", Country: "France"},
		},
	},
	{
		Name: "Japan",
		Swimmers: []model.Swimmer{
			{ID: 17, Name: "Haruto Sato", Country: "Japan"},
			{ID: 18, Name: "Yui Suzuki", Country: "Japan"},
			{ID: 19, Name: "Ren Tanaka", Country: "Japan"},
			{ID: 20, Name: "Hina Watanabe", Country: "Japan"},
		},
	},
	{
		Name: "Germany",
		Swimmers: []model.Swimmer{
			{ID: 21, Name: "Felix Muller", Country: "Germany"},
			{ID: 22, Name: "Hannah Schmidt", Country: "Germany"},
			{ID: 23, Name: "Leon Fischer", Country: "Germany"},
			{ID: 24, Name: "Mia Weber", Country: "Germany"},
		},
	},
	{
		Name: "China",
		Swimmers: []model.Swimmer{
			{ID: 25, Name: "Wei Zhang", Country: "China"},
			{ID: 26, Name: "Fang Li", Country: "China"},
			{ID: 27, Name: "Jun Wang", Country: "China"},
			{ID: 28, Name: "Mei Liu", Country: "China"},
		},
	},
	{
		Name: "Great Britain",
		Swimmers: []model.Swimmer{
			{ID: 29, Name: "Oliver Anderson", Country: "Great Britain"},
			{ID: 30, Name: "Isla Thompson", Country: "Great Britain"},
			{ID: 31, Name: "George White", Country: "Great Britain"},
			{ID: 32, Name: "Sophie Harris", Country: "Great Britain"},
		},
	},
}

var Disciplines = []model.Discipline{
	{Name: "Freestyle 50m", Style: "freestyle", Distance: 50, IsRelay: false, Legs: 1, LegDist: 50},
	{Name: "Freestyle 100m", Style: "freestyle", Distance: 100, IsRelay: false, Legs: 1, LegDist: 100},
	{Name: "Freestyle 200m", Style: "freestyle", Distance: 200, IsRelay: false, Legs: 1, LegDist: 200},
	{Name: "Freestyle 400m", Style: "freestyle", Distance: 400, IsRelay: false, Legs: 1, LegDist: 400},
	{Name: "Backstroke 100m", Style: "backstroke", Distance: 100, IsRelay: false, Legs: 1, LegDist: 100},
	{Name: "Backstroke 200m", Style: "backstroke", Distance: 200, IsRelay: false, Legs: 1, LegDist: 200},
	{Name: "Breaststroke 100m", Style: "breaststroke", Distance: 100, IsRelay: false, Legs: 1, LegDist: 100},
	{Name: "Breaststroke 200m", Style: "breaststroke", Distance: 200, IsRelay: false, Legs: 1, LegDist: 200},
	{Name: "Butterfly 100m", Style: "butterfly", Distance: 100, IsRelay: false, Legs: 1, LegDist: 100},
	{Name: "Butterfly 200m", Style: "butterfly", Distance: 200, IsRelay: false, Legs: 1, LegDist: 200},
	{Name: "Individual Medley 200m", Style: "medley", Distance: 200, IsRelay: false, Legs: 1, LegDist: 200},
	{Name: "Individual Medley 400m", Style: "medley", Distance: 400, IsRelay: false, Legs: 1, LegDist: 400},
	{Name: "4x100m Freestyle Relay", Style: "freestyle", Distance: 400, IsRelay: true, Legs: 4, LegDist: 100},
	{Name: "4x100m Medley Relay", Style: "medley", Distance: 400, IsRelay: true, Legs: 4, LegDist: 100},
	{Name: "4x200m Freestyle Relay", Style: "freestyle", Distance: 800, IsRelay: true, Legs: 4, LegDist: 200},
}

// Speed ranges per style in meters per second (min, max)
// Based on Olympic-level averages
var SpeedRanges = map[string][2]float64{
	"freestyle":    {1.78, 2.27},
	"backstroke":   {1.56, 1.92},
	"breaststroke": {1.39, 1.72},
	"butterfly":    {1.61, 2.00},
	"medley":       {1.52, 1.92},
}

// Fixed leg order for 4x100m Medley Relay per Olympic rules
var MedleyRelayOrder = []string{"backstroke", "breaststroke", "butterfly", "freestyle"}
