package medutils

const (
	minTime  = 30
	minWords = 30
)

func DetermineQuestionTime(wordCount int, media bool) int64 {
	var i int
	for wordCount > 0 {
		if inBetween(wordCount, i*minWords, (i+1)*minWords) {
			break
		}
		i++
	}
	if media {
		i++
	}

	return int64(i * 15 + minTime)
}

func inBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
