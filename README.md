# bstore
A bcrypt hashing website that you really shouldn't trust

**Creating your Cartridge**
- oc login http://youropenshiftsite
- oc new-project bstore
- oc new-build https://github.com/Jmainguy/bstore
- oc new-app --image-stream=bstore
- oc expose svc/bstore
