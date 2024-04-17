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

# DO THIS
# OUTPUT=$(SOMETHING WITH ALOT -- \
#   --OF LINES \
#   --NOT SO MUCH)

cargo run --release --manifest-path=$ManifestPath --bin collect -- \
  --task-id $TaskId \
  --leader $Leader \
  --vdaf $Vdaf \
  --collector-credential-file /home/juan/code/janus-0.7.0-prerelease-2/global-collect-config.json \
  --current-batch