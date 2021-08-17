package domain

// InputMessage The message we get from Messenger
type InputMessage struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Postback struct {
				Title   string `json:"title"`
				Payload string `json:"payload"`
			} `json:"postback"`
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Text string `json:"text"`
				Nlp  struct {
					Entities struct {
						Sentiment []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"sentiment"`
						Greetings []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"greetings"`
						Email []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"email"`
					} `json:"entities"`
					DetectedLocales []struct {
						Locale     string  `json:"locale"`
						Confidence float64 `json:"confidence"`
					} `json:"detected_locales"`
				} `json:"nlp"`
				QuickReply struct {
					Payload string `json:"payload"`
				} `json:"quick_reply"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}
