package simple_factory

import "testing"

func TestFactory_NewProduction(t *testing.T) {
	factory := Factory{}
	production1 := factory.NewProduction(1)
	production2 := factory.NewProduction(2)

	if production1.create() != "产品1" {
		t.Fatalf("产品1测试失败")
	}

	if production2.create() != "产品2" {
		t.Fatalf("产品2测试失败")
	}
}
