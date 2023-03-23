#!/bin/sh

rootDir="$(whoami) at $(date)"

mkdir "$rootDir"
mkdir "$rootDir/about_me"
mkdir "$rootDir/about_me/personal"
mkdir "$rootDir/about_me/profesional"
mkdir "$rootDir/my_friend"
mkdir "$rootDir/my_system_info"

echo http://www.facebook.com/$1 > "$rootDir/about_me/personal/facebook.txt"
echo http://www.linkedin.com/$2 > "$rootDir/about_me/profesional/linkedin.txt"

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$rootDir/my_friend/list_of_my_friend.txt"

echo "My username: $(whoami)" > "$rootDir/my_system_info/about_this_laptop.txt"
echo "with host: $(uname -a)" >> "$rootDir/my_system_info/about_this_laptop.txt"

echo "$(ping google.com -n 3)" > "$rootDir/my_system_info/internet_connection.txt"
