package gen

type InBlock interface {
	inBlock()
}

type InSection interface {
	inSection()
}

type InActions interface {
	inActions()
}

type InInput interface {
	inInput()
}

type InContext interface {
	inContext()
}
