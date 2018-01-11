package common

type CRelation struct {
	Job *CJob
	Offset int64
	Deps []*CDep
}

type CDep struct {
	Job string
	Spec string
	Offset int64
	Status uint8
}

