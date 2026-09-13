package task3

import "testing"

func TestLinkedList(t *testing.T) {
	var list LinkedList[string]
	second := list.AddLast("второй")
	list.AddFirst("первый")
	list.AddLast("четвёртый")
	third := list.InsertAfter(second, "третий")

	if list.Len() != 4 || !list.Contains("третий") || list.IndexOf("третий") != 2 {
		t.Fatal("ошибка вставки или поиска")
	}
	if list.Get(third) != "третий" {
		t.Fatal("ошибка получения значения")
	}

	list.Set(third, "новый третий")
	list.RemoveFirst()
	list.RemoveLast()
	list.Remove(second)

	if list.Len() != 1 || list.Get(third) != "новый третий" || list.IndexOf("новый третий") != 0 {
		t.Fatal("ошибка замены или удаления")
	}
	t.Log("все операции списка выполнены корректно")
}
