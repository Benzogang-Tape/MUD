package main

import (
	"testing"
)

type gameCase struct {
	step    int
	command string
	answer  string
}

var game0cases = [][]gameCase{
	[]gameCase{
		{1, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
		{2, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{3, "идти комната", "ты в своей комнате. можно пройти - коридор"},
		{4, "осмотреться", "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"},
		{5, "надеть рюкзак", "вы надели: рюкзак"},
		{6, "взять ключи", "предмет добавлен в инвентарь: ключи"},
		{7, "взять конспекты", "предмет добавлен в инвентарь: конспекты"},
		{8, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{9, "применить ключи дверь", "дверь открыта"},
		{10, "идти улица", "на улице весна. можно пройти - домой"},
	},

	[]gameCase{
		{1, "осмотреться", "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"},
		{2, "завтракать", "неизвестная команда"},
		{3, "идти комната", "нет пути в комната"}, // через стены ходить нельзя
		{4, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{5, "применить ключи дверь", "нет предмета в инвентаре - ключи"},
		{6, "идти комната", "ты в своей комнате. можно пройти - коридор"},
		{7, "осмотреться", "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"},
		{8, "взять ключи", "некуда класть"}, // надо взять рюкзак сначала
		{9, "надеть рюкзак", "вы надели: рюкзак"},
		{10, "осмотреться", "на столе: ключи, конспекты. можно пройти - коридор"}, // состояние изменилось
		{11, "взять ключи", "предмет добавлен в инвентарь: ключи"},
		{12, "взять телефон", "нет такого"},                                // неизвестный предмет
		{13, "взять ключи", "нет такого"},                                  // предмета уже нет в комнате - мы его взяли
		{14, "осмотреться", "на столе: конспекты. можно пройти - коридор"}, // состояние изменилось
		{15, "взять конспекты", "предмет добавлен в инвентарь: конспекты"},
		{16, "осмотреться", "пустая комната. можно пройти - коридор"}, // состояние изменилось
		{17, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{18, "идти кухня", "кухня, ничего интересного. можно пройти - коридор"},
		{19, "осмотреться", "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"}, // состояние изменилось
		{20, "идти коридор", "ничего интересного. можно пройти - кухня, комната, улица"},
		{21, "идти улица", "дверь закрыта"},            // условие не удовлетворено
		{22, "применить ключи дверь", "дверь открыта"}, // состояние изменилось
		{23, "применить телефон шкаф", "нет предмета в инвентаре - телефон"},
		{24, "применить ключи шкаф", "не к чему применить"}, // предмет есть, но применить его к этому нельзя
		{25, "идти улица", "на улице весна. можно пройти - домой"},
	},
}

func TestGame0(t *testing.T) {
	for caseNum, commands := range game0cases {
		initGame()
		for _, item := range commands {
			answer := handleCommand(item.command)
			if answer != item.answer {
				t.Error("case:", caseNum, item.step,
					"\n\tcmd:", item.command,
					"\n\tresult:  ", answer,
					"\n\texpected:", item.answer)
			}
		}
	}

}
