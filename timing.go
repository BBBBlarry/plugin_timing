package timing



import (
    "log"
    "time"
    "github.com/itsabot/abot/shared/datatypes"
    "github.com/itsabot/abot/shared/nlp"
    "github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
    // Abot should route messages to this plugin that contain any combination
    // of the below words. The stems of the words below are used, so you don't
    // need to include duplicates (e.g. there's no need to include both "stock"
    // and "stocks"). Everything will be lowercased as well, so there's no
    // difference between "ETF" and "etf".
    trigger := &nlp.StructuredInput{
        Commands: []string{"what"},
        Objects: []string{"date", "time"},
    }


    // Tell Abot how this plugin will respond to new conversations and follow-up
    // requests.
    fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}

    // Create the plugin.
    var err error
    pluginPath := "github.com/BBBBlarry/plugin_timing"
    p, err = plugin.New(pluginPath, trigger, fns)
    if err != nil {
        log.Fatalln("building", err)
    }

        // Add vocab handlers to the plugin
    p.Vocab = dt.NewVocab(
        dt.VocabHandler{
            Fn: kwGetTime,
            Trigger: &nlp.StructuredInput{
                Commands: []string{"what"},
                Objects: []string{"time", "date"},
            },
        },
    )
}

// Abot calls Run the first time a user interacts with a plugin
func Run(in *dt.Msg) (string, error) {
    return FollowUp(in)
}

// Abot calls FollowUp every subsequent time a user interacts with the plugin
// as long as the messages hit this plugin consecutively. As soon as Abot sends
// a message for this user to a different plugin, this plugin's Run function
// will be called the next it's triggered.  This Run/FollowUp design allows us
// to reset a plugin's state when a user changes conversations.
func FollowUp(in *dt.Msg) (string, error) {
    return p.Vocab.HandleKeywords(in), nil
}

func kwGetTime(in *dt.Msg) (resp string) {
    // Perform some lookup. We'll leave the implementation of this as an
    // exercise to reader.
    var t = "There you go: \n" + time.Now().Format(time.UnixDate)
    return t
}