# DOOMFire

**DOOMFire** is a fire effect inspired by the classic DOOM game, implemented in Go. This project showcases the power and simplicity of the Go language combined with the [Ebitengine](https://ebitengine.org), an easy-to-use gaming library with a vibrant community.

![DOOMFire](https://raw.githubusercontent.com/crgimenes/doomfire/master/doomfire.png)

## How to Run

**Note**: Some operating systems may restrict the execution of binaries downloaded from the internet for security reasons. Please consult your operating system documentation to learn how to enable DOOMFire to run.

```bash
export CGO_ENABLED=1
go run ./main.go
```

Or you can build the binary and run it:

```bash
export CGO_ENABLED=1
go build -o doomfire ./main.go
```

## How It Works

DOOMFire leverages the Ebitengine library to render a dynamic fire effect on your screen. The fire's intensity is calculated using a palette-based approach, simulating the classic flame propagation seen in the original DOOM game. The application continuously updates the fire pixels, creating a mesmerizing and nostalgic visual effect.

Then the fire is rendered on a transparent window, allowing you to see the fire effect on top of your desktop or any other application.

## Contribution

We welcome contributions to enhance DOOMFire! Whether you're fixing bugs, adding new features, or improving documentation, your help is greatly appreciated.

## License

DOOMFire is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Inspired by the classic DOOM fire effect.
- Built with the [Ebitengine](https://ebitengine.org) library.
- Special thanks to the Go community for their continuous support and contributions.

