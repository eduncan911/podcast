package podcast

// Type specifies the type of the podcast.
//
// Its values can be one of the following:
//
// Episodic (default). Specify episodic when episodes are intended to be
// consumed without any specific order. Apple Podcasts will present newest
// episodes first and display the publish date (required) of each episode.
// If organized into seasons, the newest season will be presented first
// - otherwise, episodes will be grouped by year published, newest first.
//
// For new subscribers, Apple Podcasts adds the newest, most recent episode
// in their Library.
//
// Serial. Specify serial when episodes are intended to be consumed in
// sequential order. Apple Podcasts will present the oldest episodes
// first and display the episode numbers (required) of each episode. If
// organized into seasons, the newest season will be presented first and
// <itunes:episode> numbers must be given for each episode.
//
// For new subscribers, Apple Podcasts adds the first episode to their
// Library, or the entire current season if using seasons.
const (
	Episodic Type = iota
	Serial
)

const (
	typeDefault = "Episodic"
)

// Type specifies the type of the podcast.
//
// Its values can be one of the following:
//
// Episodic (default). Specify episodic when episodes are intended to be
// consumed without any specific order. Apple Podcasts will present newest
// episodes first and display the publish date (required) of each episode.
// If organized into seasons, the newest season will be presented first
// - otherwise, episodes will be grouped by year published, newest first.
//
// For new subscribers, Apple Podcasts adds the newest, most recent episode
// in their Library.
//
// Serial. Specify serial when episodes are intended to be consumed in
// sequential order. Apple Podcasts will present the oldest episodes
// first and display the episode numbers (required) of each episode. If
// organized into seasons, the newest season will be presented first and
// <itunes:episode> numbers must be given for each episode.
//
// For new subscribers, Apple Podcasts adds the first episode to their
// Library, or the entire current season if using seasons.
type Type int

// String returns the MIME type encoding of the specified EnclosureType.
func (t Type) String() string {
	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	switch t {
	case Episodic:
		return "Episodic"
	case Serial:
		return "Serial"
	}
	return typeDefault
}
