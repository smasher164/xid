# xid

[![PkgGoDev](https://pkg.go.dev/badge/github.com/smasher164/xid)](https://pkg.go.dev/github.com/smasher164/xid)
![Test](https://github.com/smasher164/xid/workflows/Test/badge.svg)

Package xid implements validation functions for unicode identifiers,
as defined in UAX#31: https://unicode.org/reports/tr31/.
The syntax for an identifier is:

    <identifier> := <xid_start> <xid_continue>*

where `<xid_start>` and `<xid_continue>` derive from `<id_start>` and
`<id_continue>`, respectively, and check their NFKC normalized forms.