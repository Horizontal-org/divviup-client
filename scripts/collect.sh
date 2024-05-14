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

# Set variables
ManifestPath="/manifest/Cargo.toml"
TaskId="1"
Leader="https://wearehorizontal.org"
Vdaf="remove"
CredentialFile="./file.json"

############################################################
# Process the input options. Add options as needed.        #
############################################################
# Get the options
while getopts "hm:t:l:V:c:" option; do
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
      c)
        CredentialFile=$OPTARG;;
     \?) # Invalid option
         echo "Error: Invalid option"
         exit;;
   esac
done

echo $CredentialFile
set -x #echo on


cargo run --release --manifest-path=$ManifestPath --bin collect -- \
  --task-id $TaskId \
  --leader $Leader \
  --vdaf $Vdaf \
  --collector-credential-file $CredentialFile \
  --current-batch