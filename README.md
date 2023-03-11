# Openshift multiple login

## How to install
Golang 1.11 + packages  script

### Configuration script 
+ It creates users
+ Configure sudo for them
+ Remove passwords
+ Create and change ownership of the .kube folder at the home

## Use
### ./ocp - first login
Select an option and then insert the password.

```=bash
$ ocp
Ingrese opcion:
N | NAME        | URL
---------------------------------------------
0 | ocp1         | https://api.OCLUSTER1.ar:6443
1 | ocp2    | https://api.OCLUSTER2.ar:6443
2 | ocp3    | https://api.OCLUSTER3.ar:6443
3 | ocp4     | https://api.OCLUSTER4.ar:6443
4 | ocp5     | https://api.OCLUSTER5.ar:6443
5 | ocp6     | https://api.OCLUSTER6.ar:6443
6 | ocp7     | https://api.OCLUSTER7.ar:6443
---------------------------------------------
2
Password:
Login successful.

You have access to 241 projects, the list has been suppressed. You can list all projects with 'oc projects'

Using project "default".
$
```

## Build in windows

```=powershell
rsrc -manifest login.manifest -ico ocp.ico -o rsrc.syso
go get github.com/akavel/rsrc
go build
Move-Item .\ocp.exe C:\tools\ -Force
```

## Build in linux

```=bash
mv rsrc.syso asd
go get gopkg.in/yaml.v2
go build
mv ocp /usr/local/bin
mv asd rsrc.syso
```

## Copy occonfig file for linux

```=bash
cp /mnt/c/Users/user/.kube/occonfig ~/.kube/
```


## Bash script 
```
declare -r newuser="alternative"

sudo useradd -m -s /bin/bash -p $(openssl passwd -1 $newuser) $newuser

echo '$newuser ALL=(ALL) NOPASSWD:ALL' | sudo tee -a /etc/sudoers
cat 
sudo mkdir /home/$newuser/.kube && \
sudo cp /mnt/c/Users/$USER/.kube/config /home/$newuser/.kube/ && \
sudo chown -R $newuser:$newuser /home/$newuser/.kube

echo 'sudo passwd alternative'
```
