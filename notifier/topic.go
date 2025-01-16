package notifier

type topic string

const (
	UNDERDOG   topic = "UNDERDOG"
	PRIZEPICKS topic = "PRIZEPICKS"
	COMPARE    topic = "COMPARE"
	SLEEPER    topic = "SLEEPER"
)

var topicDescriptionsPlural = map[topic]string{
	UNDERDOG:   "lines added on Underdog",
	PRIZEPICKS: "lines added on PrizePicks",
	COMPARE:    "lines are available on both Underdog and PrizePicks and have differencies",
	SLEEPER:    "lines added on Sleeper",
}

var topicDescriptionsSingular = map[topic]string{
	UNDERDOG:   "line added on Underdog",
	PRIZEPICKS: "line added on PrizePicks",
	COMPARE:    "line is available on both Underdog and PrizePicks and has differencies",
	SLEEPER:    "line added on Sleeper",
}
