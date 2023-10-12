# Random image generator

This is a simple random image generator written in [Go](https://golang.org/).

Download a random image from [Lorem Picsum](https://picsum.photos/) providing a `width` and `height` parameters.

## Usage

In the root directory of the repository, run the following command:

```sh
go run .\main.go --width={width} --height={height}
```

Change the values of the `width` and `height` parameters to generate the image at the specified size.

You can find the generated image in the [`img`](./img/README.md) directory.

The generated image will have the name `{width}x{height}.png`

Where `{width}` and `{height}` are the width and height parameters provided.
