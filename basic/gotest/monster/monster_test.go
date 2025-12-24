package monster

import "testing"

func TestMonster_Store(t *testing.T) {
	monster := &Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster..Store() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.Store() 成功")
}

func TestMonster_ReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster..ReStore() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("monster.ReStore() 成功 monster=%v", *monster)
}
