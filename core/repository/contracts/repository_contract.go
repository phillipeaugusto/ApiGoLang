package contracts

type IRepository[TEntity any, TOutPut any] interface {
	IRepositoryBase[TEntity]
	Exists(entity TEntity) (error, *TOutPut)
}
