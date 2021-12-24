package mypackage

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetVacation(vacation bool) {
	e.vacation = vacation
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetVacation() bool {
	return e.vacation
}
