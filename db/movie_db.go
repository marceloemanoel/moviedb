package db

import (
	"strconv"
)

type Movie struct {
	Title        string `json:"title" bson:"title"`               // The title of the movie.
	Year         int64  `json:"year" bson:"year"`                 // The year the movie was released.
	Synopsis     string `json:"synopis" bson:"synopsis"`          // A brief summary of the movie.
	CriticScore  int    `json:"criticScore" bson:"criticScore"`   // The score given to the movie by critics.
	PeopleScore  int    `json:"peopleScore" bson:"peopleScore"`   // The score given to the movie by viewers.
	Consensus    string `json:"consensus" bson:"consensus"`       // A summary of the reviews for the movie.
	TotalReviews int64  `json:"totalReviews" bson:"totalReviews"` // The total number of reviews for the movie.
	TotalRatings int64  `json:"totalRatings" bson:"totalRatings"` // The total number of ratings for the movie.
	Type         string `json:"type" bson:"type"`                 // The type of movie (e.g. feature film, documentary, etc.).
}

const TITLE = 1
const YEAR = 2
const SYNOPSIS = 3
const CRITIC_SCORE = 4
const PEOPLE_SCORE = 5
const CONSENSUS = 6
const TOTAL_REVIEWS = 7
const TOTAL_RATINGS = 8
const TYPE = 9

func ParseMovie(record []string) *Movie {
	year, _ := strconv.ParseInt(record[YEAR], 10, 0)
	criticScore, _ := strconv.Atoi(record[CRITIC_SCORE])
	peopleScore, _ := strconv.Atoi(record[PEOPLE_SCORE])
	totalReviews, _ := strconv.ParseInt(record[TOTAL_REVIEWS], 10, 0)
	totalRatings, _ := strconv.ParseInt(record[TOTAL_RATINGS], 10, 0)

	return &Movie{
		Title:        record[TITLE],
		Year:         year,
		Synopsis:     record[SYNOPSIS],
		CriticScore:  criticScore,
		PeopleScore:  peopleScore,
		Consensus:    record[CONSENSUS],
		TotalReviews: totalReviews,
		TotalRatings: totalRatings,
		Type:         record[TYPE],
	}
}
