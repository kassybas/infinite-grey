
ffmpeg -r 1/5 -i img_%d.png -c:v libx264 -vf fps=25 -pix_fmt yuv420p out.mp4
