# uj-interface-proj
An GoLnag application to process the records and write them in the format of YAML or JSON based on the Environment varible ("FORMAT") onto a File

>**INPUT FILES**

inp.txt
Rohan,12,M,[cricket,football],5.9,50
Rohit,11,M,[football],5.8,51
Keerthi,13,F,[badminton,table tennis],5.5,45
Rohini,12,M,[tennis],5.4,44
Rakesh,12,M,[cricket],5.9,55
Vinay,12,M,[chess,carrom],5.7,52
Neha,12,F,[volleyball],5.7,40

**Key Components of the Code **

*) **Process_file()** - sport.go 
      converts the raw data of record to a compatible version of arrays of string to store them in a struct
*) **DetailWriter interface** -sport.go 
      is implemented by - DetailWriterJson struct and  DetailWriterYaml struct
      


**Internally defined Packages**
A sport.go file is placed under Sport folder 
and this is imported by main.go

**Output**
A file with all the records processed in the fomat of .yaml or .json



**External Packages used **

  *)Viper - To access Environment variables 
  
