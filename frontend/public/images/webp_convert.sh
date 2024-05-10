#!/bin/bash

webp_convert()
{
    # return if is a directory
    [[ -d "$1" ]] && return

    # return if is already a webp file
    filename=$(echo "$1" | cut -d'.' -f1)
    extension=$(echo "$1" | cut -d'.' -f2)
    [[ "$extension" == 'webp' ]] && return

    # convert to webp 
    ffmpeg -i "$1" -c:v libwebp "$filename.webp"
}

gen_small_img() {
    # return if is a directory
    [[ -d "$1" ]] && return

    # return if it is already a small image
    filename=$(echo "$1" | cut -d'.' -f1)
    extension=$(echo "$1" | cut -d'.' -f2)
    [[ "$filename" == *"_small" ]] && return

    # create a smaller version of the image
    ffmpeg -i "$1" -vf scale=20:-1 "$filename"_small."$extension"
}

for file in **/*.*; do 
    webp_convert "$file"
done
