# UwU++, the next generation esoteric language
Created by 
[Deltaphish](https://github.com/Deltaphish/UwUpp),
documentation re-written by
[blulere](https://github.com/blulere)

All language documentation is provided in the [Wiki](https://github.com/blulere/UwUpp/wiki).

**THIS IS A FORK OF DELTAPHISH'S UWU++.** I wanted to recreate it from scratch, so I didn't think about forking any of the old files over to here.

# Build Guide

Packaging the Python files into an executable is something I've not figured out how to do quite yet, and using PyInstaller causes tracebacks during execution. The Makefile provided (and this build guide) is for uwupp-go.

* In order for the Makefile to work on Windows, you'll need to run the Makefile inside of a Bash shell. The easiest way to do this is install Git with Git Bash, which works perfectly.

1. Clone the repository.
2. Type `make <os>` in a terminal. The Makefile will let you build binaries for Windows, Linux, macOS and Apple Silicon. But it's easy to build for yourself if you know how.
3. Once it's done, there will be a folder named `bin/` in `uwupp-go/`. Inside of there will be your binary.

That's it, you should now have a binary. 