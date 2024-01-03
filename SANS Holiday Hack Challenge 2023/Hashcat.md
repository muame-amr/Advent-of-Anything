# Hashcat

Links:

1. https://hashcat.net/wiki/doku.php?id=hashcat
2. https://hashcat.net/wiki/doku.php?id=example_hashes

Findings:

1. Hash type: [18200] Kerberos 5, etype 23, AS-REP

Commands used:

```bash
hashcat -a0 -m18200 hash.txt password_list.txt -w1 -n1 -u1 --force
```

Answer:
IluvC4ndyC4nes!
