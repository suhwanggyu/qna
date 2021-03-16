package thread


type Thread struct {
	topics []Topic
}

func (th *Thread) AddTopic(topic Topic) {
	th.topics = append(th.topics, topic)
}

func (th *Thread) SaveThread() {

}