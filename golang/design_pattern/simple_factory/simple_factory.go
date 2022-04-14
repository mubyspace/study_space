package simple_factory

/*
当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节
*/

type Production interface {
	create() string
}

type Factory struct{}

func (f *Factory) NewProduction(name int) Production {
	switch name {
	case 1:
		return &Production1{}
	case 2:
		return &Production2{}
	default:
		return nil
	}
}

type Production1 struct{}

func (p Production1) create() string {
	return "产品1"
}

type Production2 struct{}

func (p Production2) create() string {
	return "产品2"
}
