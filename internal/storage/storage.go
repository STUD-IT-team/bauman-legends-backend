package storage

type Storage interface {
	TextTaskStorage
	MediaTaskStorage
	TeamStorage
	ObjectStorage
	UserStorage
	SECStorage
}

type storage struct {
	TextTask  TextTaskStorage
	MediaTask MediaTaskStorage
	Team      TeamStorage
	Object    ObjectStorage
	User      UserStorage
	SEC       SECStorage
}

func NewStorage(team TeamStorage, textTask TextTaskStorage, mediaTask MediaTaskStorage, object ObjectStorage, user UserStorage, sec SECStorage) Storage {
	return &storage{
		TextTask:  textTask,
		Team:      team,
		MediaTask: mediaTask,
		Object:    object,
		User:      user,
		SEC:       sec,
	}
}
