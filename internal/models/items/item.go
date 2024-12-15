package items

type Item struct {
	Name        string
	IsNecessary bool
	IsContainer bool
	Target      Target
}

type Target interface {
	Trigger(trigger *Item) string
	Name() string
}

func (item *Item) Use(target string) string {
	if item.Target.Name() != target {
		return "не к чему применить"
	}
	return item.Target.Trigger(item)
}
