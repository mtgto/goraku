package goraku

import (
	"github.com/nlopes/slack"
	"log"
)

type Bot interface {
	// Reply to message with text.
	Reply(message *slack.MessageEvent, text string)
	// Reply to message using thread, attachments, and so on.
	ReplyWithOptions(message *slack.MessageEvent, options ...slack.MsgOption) error
	// Slack Web API client
	API() *slack.Client
	// User information of bot itself
	Me() *slack.UserDetails
	// Team which the bot belongs to
	Team() *slack.Team
}

type goraku struct {
	Bot
	api           *slack.Client
	me            *slack.UserDetails
	team          *slack.Team
	rtm           *slack.RTM
	pluginManager *pluginManager
}

// NewSlackBot returns a new slack bot.
func NewSlackBot(slackApiToken string, options ...slack.Option) *goraku {
	client := slack.New(slackApiToken, options...)
	return &goraku{
		api:           client,
		pluginManager: newPluginManager(),
	}
}

// Reply to message with text.
// Use ReplyWIthOptions for complicated message.
func (g *goraku) Reply(message *slack.MessageEvent, text string) {
	options := []slack.RTMsgOption{}
	if len(message.ThreadTimestamp) > 0 {
		options = []slack.RTMsgOption{slack.RTMsgOptionTS(message.ThreadTimestamp)}
	}
	outgoing := g.rtm.NewOutgoingMessage(text, message.Channel, options...)
	log.Printf("ThreadTimestamp: %v", message.ThreadTimestamp)
	g.rtm.SendMessage(outgoing)
}

// Reply to message using thread, attachments, and so on.
// If you want to reply like user, Use `slack.MsgOptionAsUser(true)`.
func (g *goraku) ReplyWithOptions(message *slack.MessageEvent, options ...slack.MsgOption) error {
	_, _, err := g.rtm.PostMessage(message.Channel, options...)
	return err
}

// Start RTM connection.
func (g *goraku) Start(options ...slack.RTMOption) {
	g.rtm = g.api.NewRTM(options...)
	go g.rtm.ManageConnection()

	for msg := range g.rtm.IncomingEvents {
		log.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			log.Printf("Connected: %v\n", ev)
			info := g.rtm.GetInfo()
			g.me = info.User
			g.team = info.Team
		case *slack.DisconnectedEvent:
			log.Printf("Disconnected: %v\n", ev)
			break
		case *slack.MessageEvent:
			g.pluginManager.processMessageEvent(g, ev)
			//ctx.Plugins.ExecPlugins(ctx.responseEvent(ev))
		case *slack.PresenceChangeEvent:
			log.Printf("Presence Change: %v\n", ev)
		case slack.LatencyReport:
			log.Printf("Current latency: %v\n", ev.Value)
		case *slack.RTMError:
			log.Printf("Error: %d - %s\n", ev.Code, ev.Msg)
		default:
			// Ignore other events..
			log.Printf("Unexpected: %+v\n", ev)
		}
	}
}

// Add goraku plugin to manager
func (g *goraku) AddPlugin(plugin Plugin) {
	g.pluginManager.addPlugin(plugin)
}

func (g *goraku) API() *slack.Client {
	return g.api
}

func (g *goraku) Me() *slack.UserDetails {
	return g.me
}

func (g *goraku) Team() *slack.Team {
	return g.team
}

var _ Bot = (*goraku)(nil)
