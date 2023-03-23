#!/bin/sh

rootFolder="$(whoami) at $(date)"

mkdir "$rootFolder"
mkdir "$rootFolder/about_me"
mkdir "$rootFolder/about_me/personal"
mkdir "$rootFolder/about_me/profesional"
mkdir "$rootFolder/my_friend"
mkdir "$rootFolder/my_system_info"

echo http://www.facebook.com/$1 > "$rootFolder/about_me/personal/facebook.txt"
echo http://www.linkedin.com/$2 > "$rootFolder/about_me/profesional/linkedin.txt"

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$rootFolder/my_friend/list_of_my_friend.txt"

echo "My username: $(whoami)" > "$rootFolder/my_system_info/about_this_laptop.txt"
echo "with host: $(uname -a)" >> "$rootFolder/my_system_info/about_this_laptop.txt"

echo "$(ping google.com -n 3)" > "$rootFolder/my_system_info/internet_connection.txt"
