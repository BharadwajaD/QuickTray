# QuicTray

- Single place for all my utils

- Both systray and webview want their New function to be called from main thread.. Due to which creating a seperate webview exe and integrating it with systray approach is choosen
- Keep all html files in ./html folder

# Steps to add new menu in tray

- create a seperate prog in ./cmd/ and build in the same dir
- add a switch case in ./cmd/tray/main.go along with new menu option

# Run

- ```make``` will build requried targets into ./bin dir
- make a symlink to ./bin/tray from $PATH/tray
- after logging in just run ```tray &``` in one of your terminal
- BOOM !! - all your util tools at one place
