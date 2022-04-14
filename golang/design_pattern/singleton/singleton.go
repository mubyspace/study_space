package singleton

import "sync"

/*
许多时候软件内只需要一个实例对象，此时便可用单例模式创建此对象；

某个类只能生成一个实例，该类提供了一个全局访问点供外部获取该实例，其拓展是有限多例模式
*/

var s *Singleton
var once sync.Once

type Singleton struct{}

func GetSingleton() *Singleton {
	once.Do(func() {
		s = &Singleton{}
	})

	return s
}
