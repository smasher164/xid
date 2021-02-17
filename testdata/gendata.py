# This script generates testdata for particular version of unicode.
# Used Python 3.8.7 && regex==2020.2.20 to generate cases for unicode 12.1.0.
# Used Python 3.9.1 && regex>2020.2.20 to generate cases for unicode 13.0.0.

import unicodedata
import regex
import sys
import gzip

def gendata(category):
    f = gzip.open(f'{category}{unicodedata.unidata_version}.txt.gz', 'wb')
    for r in range(sys.maxunicode+1):
        matched = bool(regex.match(f'\p{{{category}}}', chr(r)))
        if matched:
            f.write(b'T')
        else:
            f.write(b'F')

gendata('xid_start')
gendata('xid_continue')
