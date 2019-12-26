# thesis_metrics

This simple program enables to calculate difference between two images with the same size.

### Flags:
- `-o` - path to original image
- `-r` - path to second (repliacted) image
- `-g` - specify if use gray scale (default: `false`)

To run program type e.g.:
```bash
go run main.go -o ./original_image.png -r ./replicated_image.png -g
```