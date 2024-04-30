# WIP

# Properties parser 


this tool aims to help clean spring boot properties files

## Usage

### install
```bash
go build
go install
```

### Sort a Properties file
```bash
properties_parser sort -f application.properties -o application_clean.properties -g 2
```

-f source file
-o output
-g grouping level: how the tool group properties
if -g= 2 with this input file
```properties
#comment
#apzeoiazep=azeoazep
key1.sk1.ssk1=f_value1
key1.sk1.ssk2=f_value2
key2.sk1.ssk1=f_value6
key1.sk2.ssk2=f_value3
#amzelazeopi
key1.sk2.ssk1=f_value4
```
the output file will be 
```properties
key1.sk1.ssk1=f_value1
key1.sk1.ssk2=f_value2

key1.sk2.ssk1=f_value4
key1.sk2.ssk2=f_value3

key2.sk1.ssk1=f_value6
```
