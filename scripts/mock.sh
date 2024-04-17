#!/bin/bash

#!/bin/bash
############################################################
# Help                                                     #
############################################################
Help()
{
   # Display Help
   echo "Collection"
   echo
   echo "Syntax: collect [-g|h|v|V]"
   echo "options:"
   echo "h     Print this Help."
   echo "v     Verbose mode."
   echo
}

############################################################
############################################################
# Main program                                             #
############################################################
############################################################


############################################################
# Process the input options. Add options as needed.        #
############################################################
# Get the options
while getopts "hm:t:l:V:" option; do
   case $option in
      h) # display Help
         Help
         exit;;
      m) 
         ManifestPath=$OPTARG;;
      t)
        TaskId=$OPTARG;;
      l)
        Leader=$OPTARG;;
      V)
        Vdaf=$OPTARG;;
     \?) # Invalid option
         echo "Error: Invalid option"
         exit;;
   esac
done


# TODO make collector-credential env variable

set -x #echo on

if [ $Vdaf == "count" ] 
then
   echo "Count: 200"
fi

if [ $Vdaf == "sum" ] 
then
   echo "Sum: 230"
fi

if [ $Vdaf == "histogram" ] 
then
   echo "Histogram: [8, 20, 4]"
fi