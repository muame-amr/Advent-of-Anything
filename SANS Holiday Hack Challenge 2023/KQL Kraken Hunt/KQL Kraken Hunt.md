# KQL Kraken Hunt

## Questions

### Introduction

1. How many craftperson elf works from laptop ?

   ```kql
   Employees
   | where role == "Craftsperson Elf"
   | where hostname has "LAPTOP"
   | count
   ```

   Answer: 25

### Case 1

Malicious link: http://madelvesnorthpole.org/published/search/MonthlyInvoiceForReindeerFood.docx

1. What is the email address of the employee who received this phishing email?
2. What is the email address that was used to send this spear phishing email?
3. What was the subject line used in the spear phishing email?

   ```kql
   Email
   | where link == "http://madelvesnorthpole.org/published/search/MonthlyInvoiceForReindeerFood.docx"
   | where recipient in (
       Employees
       | where ip_addr in (
           OutboundNetworkEvents
           | where url has "http://madelvesnorthpole.org/published/search/MonthlyInvoiceForReindeerFood.docx"
           | project src_ip
           )
       | project email_addr)
   | project recipient, sender, subject
   ```

   Answer:

   | recipient                                        | sender             | subject                                        |
   | ------------------------------------------------ | ------------------ | ---------------------------------------------- |
   | alabaster_snowball@santaworkshopgeeseislands.org | cwombley@gmail.com | [EXTERNAL] Invoice foir reindeer food past due |

### Case 2

1. What is the role of our victim in the organization?
2. What is the hostname of the victim's machine?
3. What is the source IP linked to the victim?

   ```kql
   Employees
   | where email_addr == "alabaster_snowball@santaworkshopgeeseislands.org"
   | project role, hostname, ip_addr
   ```

   Answer:

   | role     | hostname     | ip_addr   |
   | -------- | ------------ | --------- |
   | Head Elf | Y1US-DESKTOP | 10.10.0.4 |

### Case 3

1. What time did Alabaster click on the malicious link? Make sure to copy the exact timestamp from the logs!

   ```kql
   OutboundNetworkEvents
   | where url has "http://madelvesnorthpole.org/published/search/MonthlyInvoiceForReindeerFood.docx"
   | project timestamp;
   ```

   Answer: 2023-12-02T10:12:42Z

2. What file is dropped to Alabaster's machine shortly after he downloads the malicious file?

   ```kql
   FileCreationEvents
   | where username == "alsnowball"
   | where timestamp between(datetime(2023-12-02T10:14:00Z) .. datetime(2023-12-02T10:15:00Z))
   | project filename
   ```

   Answer: giftwrap.exe

### Case 4

1. The attacker created an reverse tunnel connection with the compromised machine. What IP was the connection forwarded to?

   ```kql
   ProcessEvents
   | where username == "alsnowball"
   | where timestamp between(startofday(datetime(2023-12-02)) .. endofday(datetime(2023-12-03)))
   | where process_commandline has "bind"
   ```

   You will find forwarded IP in `process_commandline` column

   Answer: 113.37.9.17:22

2. What is the timestamp when the attackers enumerated network shares on the machine?

   ```kql
   ProcessEvents
   | where username == "alsnowball"
   | where timestamp between(startofday(datetime(2023-12-02)) .. endofday(datetime(2023-12-03)))
   ```

   Find when the attacker used `net share` in `process_commandline`

   Answer: 2023-12-02T16:51:44Z

3. What was the hostname of the system the attacker moved laterally to?

   ```kql
   ProcessEvents
   | where process_commandline has "net use"
   ```

   Answer: NorthPolefileshare

### Case 5

1. When was the attacker's first base64 encoded PowerShell command executed on Alabaster's machine?

   ```kql
   ProcessEvents
   | where process_name has "powershell.exe"
   | where hostname == "Y1US-DESKTOP";
   ```

   Find base64 encoded command in `process_commandline` column. Decode using any base64 decoding tools

   Answer: 2023-12-24T16:07:47Z

2. What was the name of the file the attacker copied from the fileshare? (This might require some additional decoding)

   Answer: NaughtNiceList.docx

3. The attacker has likely exfiltrated data from the file share. What domain name was the data exfiltrated to?

   Answer: giftbox.com

### Case 6

```kql
  ProcessEvents
  | where process_name has_any("powershell.exe", "cmd.exe")
  | where hostname == "Y1US-DESKTOP";
```

1. What is the name of the executable the attackers used in the final malicious command?

   Answer: downwithsanta.exe

2. What was the command line flag used alongside this executable?

   Answer: --wipeall
