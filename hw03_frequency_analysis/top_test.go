package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

*/

var text0 = `они, слава богу, абсолютно нормальные люди. Уж от кого-кого,
	а от них никак нельзя было ожидать, чтобы они попали в какую-нибудь
	странную или загадочную ситуацию. Мистер и миссис Дурсль весьма неодобрительно
	относились к любым странностям, загадкам и прочей ерунде.
	Мистер Дурсль возглавлял фирму под названием «Граннингс»,
	которая специализировалась на производстве дрелей.
	Это был полный мужчина с очень пышными усами и очень короткой шеей.
	Что же касается миссис Дурсль, она была тощей блондинкой с шеей почти вдвое длиннее,
	чем положено при ее росте. Однако этот недостаток пришелся ей весьма кстати,
	поскольку большую часть времени миссис Дурсль следила за соседями и подслушивала их разговоры.
	А с такой шеей, как у нее, было очень удобно заглядывать за чужие заборы.
	У мистера и миссис Дурсль был маленький сын по имени Дадли,
	и, по их мнению, он был самым чудесным ребенком на свете.`

/*
func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}
*/

func TestGarryPotterTop10(t *testing.T) {
	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"и",      // 5
			"Дурсль", // 4
			"миссис", // 4
			"был",    // 3
			"очень",  // 3
			"с",      // 3
			"Мистер", // 2
			"было",   // 2
			"весьма", // 2
			"за",     // 2
		}
		require.Equal(t, expected, Top10(text0))
	})
}
