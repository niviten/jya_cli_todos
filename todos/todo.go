package todos

type Todo struct {
    ID int64
    Name string
    Description string
    CreatedDate string
    DueDate string
    IsCompleted bool
    CompletedDate string
    Priority int
    Tags []int64
    Project int64
}

func New() *Todo {
    return &Todo{ IsCompleted: false, Priority: 5 }
}
