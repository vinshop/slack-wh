package slack_wh

type Block interface {
	block()
}

type Element interface {
	element()
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
