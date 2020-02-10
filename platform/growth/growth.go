package growth

//opcional para mas control
//por medio de las interfaces implicitas, GetAll se refiere a Repo.GetAll
type Getter interface {
	GetAll() []Item
}

type Setter interface {
	Add(item Item)
}

type Item struct {
	Value   string `json:"value"`
	Created string `json:"created"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	//lo inicializa vacio
	return &Repo{
		Items: []Item{},
	}

	//lo inicializa en null
	//return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
