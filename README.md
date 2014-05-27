# MUMGO

GO package to build mumble client bot.

NOTE: this is still very much work in progress ... use it at ur own risk.


## Usage

Example usage:

    package "main"

    import "github.com/handymic/mumgo"

    func main() {

      // Configure
      config = mumgo.Config(...)

      // Initializes a bot
      bot := mumgo.Connect(config)

      // Upon receiving text message
      bot.OnText(func(text mumble.Text){
        // ...
      })

      // Upon receiving audio message
      bot.OnAudio(func(audio mumble.Audio){
        // ...
      })

    }


## License

See LICENSE

