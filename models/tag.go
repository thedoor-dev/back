package models

type Tag struct {
	ID   int64  `db:"tid"  json:"tid"`
	PID  int64  `db:"pid"  json:"pid"`
	Name string `db:"name" json:"name"`
}

type TagArr []Tag

// func (t Tags) Marshal() ([]byte, error) {

// 	log.Println(t)
// 	var res []byte
// 	var r []byte
// 	var err error
// 	res = append(res, byte('['))
// 	for _, v := range t {
// 		r, err = json.Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		res = append(res, r...)
// 	}
// 	// res = append(res, ']')
// 	return res, err
// }
