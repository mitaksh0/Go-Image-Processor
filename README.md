# Welcome to Image Processing Tool

It works as backend service instead of API, you would need to pass arguments in commandline to run this project. 

# How It Works:
Add original (to be changed) images at "./img" 
Add any or all arguments launch.json or pass it in cli command: "scale:upscale:1:3|grayscale|compress:80" (without quotes in cli OR with quotes in launch.json, "|" separates each process). This will do the following:
1. Upscale image using fastest scaling function(nearest neighbor) with 2x.
(To downscale replace "upscale" with "downscale"; To use other scaling techniques : 1:Nearest Neighbor, 2:BiLinear, 3:Catmull-Rom kernel but it will keep slowing down as the number increase as each will create better quality and previous; 3 represents 3x (Scaling factor); 80 is the compression level(100 is minimum compression with best quality and vice versa is 1.))
2. After passing the appropriate commands, processed image will be saved on "./processed_img" directory and will be named same as image plus the transformations applied.

# END
