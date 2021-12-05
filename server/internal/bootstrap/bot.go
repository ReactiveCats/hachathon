package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"github.com/tucnak/telebot"
	"server/external/fsm"
	"server/internal/domain"
	"strconv"
	"strings"
	"time"
)

type TelegramBOT struct {
	Bot  *telebot.Bot
	deps Dependencies
}

func NewTGBot(deps Dependencies) (TelegramBOT, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  deps.Config.TelegramTokenBot,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return TelegramBOT{}, errors.New("wrong token")
	}

	return TelegramBOT{Bot: bot, deps: deps}, nil
}

var states = map[int64]string{}
var argsMap = map[int64][]string{}

func (b TelegramBOT) Start() {
	b.Bot.Handle(telebot.OnText, func(m *telebot.Message) {
		state, ok := states[m.Chat.ID]
		if !ok || m.Text == "/start" || m.Text == "/tasks" {
			state = "-"
		}
		args, ok := argsMap[m.Chat.ID]
		if !ok {
			args = []string{}
		}

		fsm := b.NewFSM(b.Bot, m)
		state, args, err := fsm.Handle(map[string]interface{}{}, state, args, m.Text)
		if err != nil {
			switch err.(type) {
			case WrongInputError:
				b.Bot.Reply(m, err.Error())
				return
			}
			fmt.Println(err)
		}
		states[m.Chat.ID] = state
		argsMap[m.Chat.ID] = args
	})

	b.Bot.Start()
}

func (b TelegramBOT) NewFSM(bot *telebot.Bot, message *telebot.Message) *fsm.BotFSM {
	botFSM := fsm.NewBotFSM()

	botFSM.RegisterMenu("-",
		fsm.NewTransition(func(ctx *fsm.Context) (fsm.Arguments, error) {
			bot.Reply(message, "Вітаю! Я маю такі команди:\n"+
				"/new_task - створити нове завдання\n"+
				"/tasks - подивитися існуючи завдання")
			return fsm.Arguments{}, nil
		}, fsm.StaticText(""), fsm.StaticKeyboard()),
		fsm.Transitions{
			"/new_task":       "type_task_title",
			"/tasks":          "tasks",
			fsm.WildcardToken: "-",
		})

	botFSM.RegisterMenu("tasks",
		fsm.NewTransition(func(ctx *fsm.Context) (fsm.Arguments, error) {
			user := b.auth(message.Sender)
			sb := strings.Builder{}

			priority := true
			tasks, err := b.deps.TaskService.Fetch(context.Background(), domain.GetTaskDTO{
				UserID:   user.ID,
				Priority: &priority,
			})
			if err != nil {
				return fsm.Arguments{}, err
			}

			if len(tasks) == 0 {
				bot.Reply(message, "Поки що завдань немає, але все попереду!")
				return fsm.Arguments{}, err
			}

			for i, task := range tasks {
				sb.WriteString(fmt.Sprintf("%d) [%.2f] %s\n", i+1, *task.FinalPriority, task.Title))
			}
			bot.Reply(message, sb.String())
			return fsm.Arguments{}, nil
		}, fsm.StaticText(""), fsm.StaticKeyboard(), fsm.GoForwardSilent(true)),
		fsm.Transitions{
			fsm.WildcardToken: "-",
		})

	botFSM.RegisterMenu("type_task_title",
		fsm.NewTransition(func(ctx *fsm.Context) (fsm.Arguments, error) {
			bot.Reply(message, "Введіть назву завдання")
			return fsm.Arguments{}, nil
		}, fsm.StaticText(""), fsm.StaticKeyboard()),
		fsm.Transitions{
			fsm.WildcardToken: "type_task_parameters_ask",
		})

	botFSM.RegisterMenu("type_task_parameters_ask",
		fsm.NewTransition(func(ctx *fsm.Context) (fsm.Arguments, error) {
			bot.Reply(message, "Введіть важливість та терміновість завдання, від 1 до 3 (например, \"2 3\")")
			return fsm.Arguments{message.Text}, nil
		}, fsm.StaticText(""), fsm.StaticKeyboard()),
		fsm.Transitions{
			fsm.WildcardToken: "type_task_parameters",
		})

	botFSM.RegisterMenu("type_task_parameters",
		fsm.NewTransition(func(ctx *fsm.Context) (fsm.Arguments, error) {
			user := b.auth(message.Sender)

			fields := strings.Fields(message.Text)
			if len(fields) != 2 {
				return ctx.Arguments, WrongInputError{Text: "Введіть важливість та терміновість завдання, від 1 до 3 (например, \"2 3\")"}
			}

			_, err := strconv.Atoi(fields[0])
			if err != nil || (fields[0] != "1" && fields[0] != "2" && fields[0] != "3") {
				return ctx.Arguments, WrongInputError{Text: "Введіть важливість та терміновість завдання, від 1 до 3 (например, \"2 3\")"}
			}

			_, err = strconv.Atoi(fields[1])
			if err != nil || (fields[1] != "1" && fields[1] != "2" && fields[1] != "3") {
				return ctx.Arguments, WrongInputError{Text: "Введіть важливість та терміновість завдання, від 1 до 3 (например, \"2 3\")"}
			}
			ctx.Arguments = append(ctx.Arguments, fields[0], fields[1])

			// 1
			// 2
			// 3

			// 2
			// 4
			// 6

			// 3
			// 5
			// 7

			// todo save task to service
			t := time.Now().Add(24 * time.Hour)
			est := 60 * 60 * 5
			task, _, err := b.deps.TaskService.Create(context.Background(), domain.CreateTaskDTO{
				UserID:     user.ID,
				Title:      ctx.Arg(0),
				Deadline:   &t,
				Estimated:  &est,
				Importance: ctx.ArgInt(1)*2 + 1,
				Urgency:    ctx.ArgInt(2)*2 + 1,
			})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			bot.Reply(message, fmt.Sprintf("Нове завдання \"%s\" успішно створене!", task.Title))

			return fsm.Arguments{}, nil
			//return fsm.Arguments{ctx.Arg(0), fields[0], fields[1]}, nil
		}, fsm.StaticText(""), fsm.StaticKeyboard(), fsm.GoForward(true)),
		fsm.Transitions{
			fsm.WildcardToken: "-",
		})

	return botFSM
}

func (b TelegramBOT) auth(user *telebot.User) *domain.User {
	username := user.Username
	if username == "" {
		username = strconv.Itoa(user.ID)
	}

	_, _ = b.deps.UserService.Signup(context.Background(), user.Username)

	u, err := b.deps.UserService.ByUsername(context.Background(), username)
	if err != nil {
		return nil
	}

	return u
}

type WrongInputError struct {
	Text string
}

func (o WrongInputError) Error() string {
	return o.Text
}
