# Print Hello World under a chroot directory using Go

Prerequisites :
- [Install Go](https://go.dev/doc/install)


**Execute all the following commands as root or using sudo** <br>
**Execute all the following from this current directory**

Creating the chroot directory
```bash
sudo mkdir /tmp/chroot
```

Adding /dev/null in the chroot directory
```bash
sudo mkdir -p /tmp/chroot/dev
sudo mknod -m 666 /tmp/chroot/dev/null c 1 3
sudo mknod -m 666 /tmp/chroot/dev/zero c 1 5
sudo mknod -m 666 /tmp/chroot/dev/random c 1 8
sudo mknod -m 666 /tmp/chroot/dev/urandom c 1 9
```

Adding bash with his libraries in the chroot directory
```bash
# Creating necessary directories
sudo mkdir -p /tmp/chroot/{bin,lib,lib64}

# Copying bash binary
sudo cp /bin/bash /tmp/chroot/bin/

# Copying necessary libraries
for lib in $(ldd /bin/bash | grep -o '/[^ ]*'); do
    sudo mkdir -p /tmp/chroot$(dirname $lib)
    sudo cp $lib /tmp/chroot$lib
done
```

Copying the hello.sh script under the chroot directory and making it executable
```bash
sudo cp hello.sh /tmp/chroot/
sudo chmod +x /tmp/chroot/hello.sh
```

Execute the go program
```bash
go run main.go
```


The result should be :
```bash
Hello World from plonkc :)
```