package maps

// var m map[string]model.User

// func New() {
// 	m = make(map[string]model.User)
// }

// func Add(key string, value model.User) {
// 	m[key] = value
// }

// func Get(key string) model.User {
// 	return m[key]
// }

// func Delete(key string) {
// 	delete(m, key)
// }

type MyMap map[string]string

// func New
func New() MyMap {
	return make(MyMap)
}

// func Add
func (m MyMap) Add(key string, value string) {
	m[key] = value
}

// func Get
func (m MyMap) Get(key string) string {
	return m[key]
}

// func Delete une clée d'une map
func (m MyMap) Delete(key string) {
	delete(m, key)
}

// func delete map
