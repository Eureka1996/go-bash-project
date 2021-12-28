package monster_test

import (
	monster "project2/src/go_code/project2/unitTest"
	"testing"
)

func TestMonsterStore(t *testing.T) {
	monster := &monster.Monster{
		Name:  "wufuqiang",
		Age:   20,
		Skill: "hahaha",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v,实际为=%v", true, res)
	}
	t.Logf("monster.Store测试成功。")
}

func TestMonsterReStore(t *testing.T) {
	monster := &monster.Monster{}

	res := monster.ReStore()
	if !res || monster.Age != 19 {
		t.Fatalf("monster.ReStore() 错误，希望为:result=%v,age=%v,实际为:result=%v,age=%v",
			true, 19, res, monster.Age)
	}
	t.Logf("monster.Store测试成功。")
}
