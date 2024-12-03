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

type MyMap map[string]interface{}

// func New
func New() MyMap {
	return make(MyMap)
}

// func Add
func (m MyMap) Add(key string, value interface{}) {
	m[key] = value
}

// func Get
func (m MyMap) Get(key string) interface{} {
	return m[key]
}

// func Delete une clée d'une map
func (m MyMap) Delete(key string) {
	delete(m, key)
}

// func delete map
