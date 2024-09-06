package storage

type Storage interface {
	TextTaskStorage
	MediaTaskStorage
	TeamStorage
	ObjectStorage
	UserStorage
}

type storage struct {
	TextTask  TextTaskStorage
	MediaTask MediaTaskStorage
	Team      TeamStorage
	Object    ObjectStorage
	User      UserStorage
}

func NewStorage(team TeamStorage, textTask TextTaskStorage, mediaTask MediaTaskStorage, object ObjectStorage, user UserStorage) Storage {
	return &storage{
		TextTask:  textTask,
		Team:      team,
		MediaTask: mediaTask,
		Object:    object,
		User:      user,
	}
}
