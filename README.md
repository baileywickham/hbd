# hbd

For doing the absolute minimum on someone's birthday.

The program runs through a contact list and sends a hbd text to everyone if it is their birthday, including a prebuilt "nice" message. Built for fun as a joke.


Example message:
> hbd Bailey Wickham! you are very Neat!


## Use
Export contacts as `.vcf` file, this is supported by Apple and Google.

This is built on Go, so compile with `go build .` and run with:
```bash
./hbd <contact.vcf>
```
