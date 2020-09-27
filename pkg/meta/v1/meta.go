package v1

type Resource interface {
	GetID() uint64
	GetCreateTime() JsonTime
	GetUpdateTime() JsonTime
}

func (meta *ObjectMeta) GetID() uint64           { return meta.ID }
func (meta *ObjectMeta) GetCreateTime() JsonTime { return meta.CreatedAt }
func (meta *ObjectMeta) GetUpdateTime() JsonTime { return meta.UpdatedAt }
