# spotify-remote

This little cli tool allows you to control playback of a Spotify client from another PC.
The goal is to have a low-resource remote to be able to control the Spotify client that is running on another device.

For now, it works by using a *keylogger* to detect the play/pause/next keystrokes (with no ill intent).  
It only communicates with the Spotify API.

Future TODOs:
- [ ] Show currently playing song
- [ ] Store and reuse authentication token
- [ ] Use a CLI audio player to take ownership of the playback controls instead of using a keylogger.

Contributions welcome!

License: MIT
