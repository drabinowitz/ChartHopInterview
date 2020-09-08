package dao

// DAO interface any datalayer must implement
type DAO interface {
	Read(userID, recordID string) (map[string]interface{}, error)
	List(userID string) ([]map[string]interface{}, error)
	Create(record interface{}) error
	Update(record interface{}) error
	Remove(userID, recordID string) error
}
