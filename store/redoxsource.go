package store

import "github.com/jinzhu/gorm"

//GetRedoxSourceBySourceID GetRedoxSourceBySourceID
func (s *Store) GetRedoxSourceBySourceID(sourceID string) *gorm.DB {

	return s.Database.Where(`"redoxId" = ?`, sourceID)
}
