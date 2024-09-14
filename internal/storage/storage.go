package storage

type Storage interface {
	TextTaskStorage
	MediaTaskStorage
	TeamStorage
	ObjectStorage
	UserStorage
	MasterClassStorage
}

type storage struct {
	TextTask    TextTaskStorage
	MediaTask   MediaTaskStorage
	Team        TeamStorage
	Object      ObjectStorage
	User        UserStorage
	MasterClass MasterClassStorage
}

func NewStorage(team TeamStorage, textTask TextTaskStorage, mediaTask MediaTaskStorage, object ObjectStorage, user UserStorage, masterClass MasterClassStorage) Storage {
	return &storage{
		TextTask:    textTask,
		Team:        team,
		MediaTask:   mediaTask,
		Object:      object,
		User:        user,
		MasterClass: masterClass,
	}
}
