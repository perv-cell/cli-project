# PELP
# path-helper for system 

Environment variable helper. When writing a project for testing or debugging, you often use a config file intended for local development. Some data from this config file should be stored secretly on a remote machine. To do this, we can test this data on our local machine by adding and reading it from the PATH. This cli application will speed up your process.

How to use:

1. Clone this repository.

2. Go to the project root and cd path-helper.

3. Run go build -o directory where you want to save the binary.

4. Add the path to this binary to the PATH:
On Windows: Go to Environment Variables and add it to the PATH.
On Linux, find the bash.bashrc file at /etc/bash.bashrc (as root) or ~/.bashrc (as user) and add the line:
export PATH="$PATH:path to your binary folder"

Run the source command "with the path to the folder where the bashrc file was located.

Go to the command line and check the application's operation with: pelp --help. If everything is OK, congratulations! Working with Path will become much easier.
