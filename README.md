# Downloader
CLI based downloader utilising the concurrency feature of Golang. 

- Takes the download URL of the file from the user.
- Splits the files into sections which then downloads all the chunks asynchronously.
- The temp files are then merged to the filename provided by the user.
