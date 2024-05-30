package event

import (
	"fmt"
	"math/rand"
)

// 主题接口
type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

// 观察者接口
type Observer interface {
	Update(string)
}

// 具体主题，即新闻发布者
type NewsPublisher struct {
	observers []Observer
}

func NewNewsPublisher() *NewsPublisher {
	return &NewsPublisher{}
}

func (np *NewsPublisher) RegisterObserver(o Observer) {
	np.observers = append(np.observers, o)
}

func (np *NewsPublisher) RemoveObserver(o Observer) {
	for i, observer := range np.observers {
		if observer == o {
			np.observers = append(np.observers[:i], np.observers[i+1:]...)
			break
		}
	}
}

func (np *NewsPublisher) NotifyObservers() {
	// 模拟发布新闻
	news := generateNews()
	fmt.Printf("NewsPublisher: News - %s\n", news)

	// 通知所有订阅者
	for _, observer := range np.observers {
		observer.Update(news)
	}
}

// 具体观察者，即新闻订阅者
type NewsSubscriber struct {
	name string
}

func NewNewsSubscriber(name string) *NewsSubscriber {
	return &NewsSubscriber{name}
}

func (ns *NewsSubscriber) Update(news string) {
	fmt.Printf("%s received news: %s\n", ns.name, news)
}

// 生成随机新闻
func generateNews() string {
	news := []string{
		"Breaking: Major Event Happens",
		"Technology Advancements Announced",
		"Stock Market Update",
		"Weather Forecast",
	}
	return news[randInt(0, len(news))]
}

// 随机整数生成器
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
