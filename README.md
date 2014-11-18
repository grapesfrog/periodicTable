periodicTable
=============

Learning Go while hopefully producing some useful code along the way. 
I am using a Periodic table of the elements  as the Dataset but I cannot recall where I obtained the orignal csv. If you recognise the following header let me know so I can acknowledge:

"Z", "Symbol", "Name", "A", "N", "Period", "IUPAC_Group", "Old_IUPAC", "CAS_Group", "Category", "Standard_Atomic_Weight", "Last _Digit", "Ionisation_eV", "Normal_State", "Density_Kg_m3_20°C", "Melting_point_°C", "Melting_point_K", "Boiling_point_°C", "Boiling_point_K", "Year_of_discovery", "Discoverer"

Version 1:

parse-csv.go

The first program  takes a csv of the periodic elements and converts to a flat JSON file. 

Version 2:

In version 2 the code is refactored so that the CreateNewFile and AppendToFile functions are removed from the main program (which is now called flat-json.go ) and split out into a separate package called parser (I know I am rubbish at naming anything!) The main program calls the package which contains the functions and the program works as before.

This reconfiguration is so that the code can be set up following best practise in terms of using your  [workspace](http://golang.org/doc/code.html#Workspaces)

**To build and run**

Build and install package

Assuming you have cloned into your workspace

cd into the parsers folder and build and install the parsers package

`go build `

`go install`

Then build the flat-json program 

`go build flat-json`

To run 

`./flat-json -ouput=name-of-generated-output-file -input=name-of-csv-file-to-read` 



**Testing**

To run the tests type 

`go test`


**wikiparse.go** 

I needed a bigger data set for some vraious tests I was doing so I extracted data from wikipedia using the category "Chemical elements" and removing the category and random pages that I would have got via this prior to doanloading from wikipedia

This code should work with any collection of articles that have been downloaded as an XML extract from wikipedia just plug in your own XML file.
