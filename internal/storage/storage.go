package storage

type Storage interface {
	TextTaskStorage
	MediaTaskStorage
	TeamStorage
	ObjectStorage
}

type storage struct {
	TextTask  TextTaskStorage
	MediaTask MediaTaskStorage
	Team      TeamStorage
	Object    ObjectStorage
}

func NewStorage(team TeamStorage, textTask TextTaskStorage, mediaTask MediaTaskStorage, object ObjectStorage) Storage {
	return &storage{
		TextTask:  textTask,
		Team:      team,
		MediaTask: mediaTask,
		Object:    object,
	}
}
