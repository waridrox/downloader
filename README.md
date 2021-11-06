# Downloader
CLI based downloader utilising the concurrency feature of Golang. 

- Takes the download URL of the file from the user.
- Splits the files into sections which then downloads all the chunks asynchronously.
- The temp files are then merged to the filename provided by the user.

Demo: 

https://user-images.githubusercontent.com/58583793/140624089-01220d26-2b37-4679-b220-13d19914f850.mp4
