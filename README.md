# Go Asteroids

A classic Asteroids arcade game implementation written in Go using the Ebitengine game library. This project demonstrates 2D game development fundamentals including sprite rendering, collision detection, game loops, and scene management.

## Screenshot

![Game Screenshot](assets/images/player.png)

## Features

- Classic Asteroids gameplay mechanics
- Smooth 60 FPS gameplay
- Collision detection system
- Multiple game scenes (title, game, game over)
- Sound effects and background music
- WebAssembly support for browser deployment
- Cross-platform compatibility (Windows, macOS, Linux)

## Game Controls

- **Arrow Keys**: Move the spaceship
- **Spacebar**: Fire laser
- **H**: Activate hyperspace (teleport to random location)
- **S**: Toggle shield
- **P**: Pause game

## Gameplay

- Destroy asteroids to earn points
- Large asteroids break into smaller pieces when destroyed
- Avoid collisions with asteroids and alien ships
- Collect power-ups for enhanced abilities
- Survive as long as possible to achieve high scores

## Technical Details

### Architecture

The game follows a scene-based architecture with separate components for:
- **Scene Manager**: Handles transitions between game states
- **Game Objects**: Player, asteroids, lasers, aliens, and power-ups
- **Collision System**: Efficient collision detection between game entities
- **Audio System**: Sound effects and background music management
- **Input Handling**: Keyboard and mouse input processing

### Dependencies

- **Ebitengine v2**: 2D game library for Go
- **Go modules**: Dependency management
- **Embed**: Asset embedding for distribution

### Build Targets

- **Native**: Windows, macOS, Linux executables
- **WebAssembly**: Browser-based gameplay
- **Mobile**: iOS and Android support via gomobile

## Installation

### Prerequisites

- Go 1.21 or later
- Git

### Building from Source

```bash
# Clone the repository
git clone https://github.com/faheem047/go-asteroids.git
cd go-asteroids

# Install dependencies
go mod download

# Build for your platform
go build -o go-asteroids main.go

# Run the game
./go-asteroids
```

### WebAssembly Build

```bash
# Build for web
GOOS=js GOARCH=wasm go build -o wasm/goasteroids.wasm main.go

# Serve the wasm directory with a web server
cd wasm
python -m http.server 8000
# Open http://localhost:8000 in your browser
```

### Cross-Platform Builds

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o go-asteroids.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o go-asteroids main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o go-asteroids main.go
```

## Project Structure

```
go-asteroids/
├── assets/                 # Game assets (images, audio, fonts)
│   ├── images/            # Sprites and textures
│   ├── audio/             # Sound effects and music
│   └── fonts/             # Game fonts
├── goasteroids/           # Main game source code
│   ├── game.go            # Core game loop and state
│   ├── player.go          # Player spaceship logic
│   ├── meteor.go          # Asteroid behavior
│   ├── alien.go           # Enemy AI and behavior
│   ├── laser.go           # Projectile system
│   ├── collision.go       # Collision detection
│   └── scenes/            # Game scene implementations
├── wasm/                  # WebAssembly build files
├── main.go                # Entry point
├── go.mod                 # Go module definition
└── README.md              # This file
```

## Development

### Adding New Features

1. **Game Objects**: Create new structs implementing the GameObject interface
2. **Scenes**: Add new scene types to the scene manager
3. **Assets**: Place new sprites/audio in the assets directory
4. **Collisions**: Register new collision types in the collision system

### Code Style

- Follow Go standard formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for complex game logic
- Keep functions focused and concise

### Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...
```

## Performance Considerations

- **Sprite Batching**: Group similar draw operations
- **Object Pooling**: Reuse game objects to reduce allocations
- **Spatial Partitioning**: Optimize collision detection for large numbers of objects
- **Asset Management**: Efficiently load and cache game resources

## Deployment

### Desktop Applications

The game can be packaged as standalone applications for distribution:

- **Windows**: Create installer with tools like NSIS
- **macOS**: Build .app bundle and create .dmg
- **Linux**: Package for distribution repositories

### Web Deployment

Deploy to any static hosting service:
- GitHub Pages
- Netlify
- Vercel
- AWS S3

### Mobile Deployment

Use gomobile for iOS and Android builds:

```bash
# Install gomobile
go install golang.org/x/mobile/cmd/gomobile@latest

# Build for Android
gomobile build -target=android

# Build for iOS
gomobile build -target=ios
```

## Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

This project was inspired by the excellent tutorial "Making Games in Go for Absolute Beginners" by Miłosz Smółka at [Three Dots Labs](https://threedots.tech/post/making-games-in-go/). The tutorial provided valuable insights into game development fundamentals using Go and Ebitengine.

## Resources

- [Ebitengine Documentation](https://ebitengine.org/)
- [Go Game Development Community](https://github.com/topics/go-game)
- [Game Development Patterns](https://gameprogrammingpatterns.com/)
- [2D Game Art Resources](https://kenney.nl/assets)

## Support

If you encounter any issues or have questions:

1. Check the existing issues on GitHub
2. Search the documentation
3. Create a new issue with detailed information
4. Join the Go game development community discussions

---

**Note**: This is a hobby project created for learning and entertainment purposes. The game mechanics and implementation are simplified versions of the classic Asteroids arcade game.
