#! /bin/bash

# Note: This is for installing all neccessary tools for this server
#It will run down a list of config file entries, then for each entry, 
# see if it can run the command in the yaml file

#Get DLF Load time
DLF_LOAD_DTTM=$(date '+%Y-%m-%d_%H-%m-%S')

#File location:
file_location="/root/startUpCronJob/logging"
file_name="installlinux_${DLF_LOAD_DTTM}.log"

#This is the log file location:
LOG_FILE="${file_location}/${file_name}"

#Has a list of all the software/libraries we will install
library_install_list="/root/startUpCronJob/config/library_install_list.txt"

#Get listed tables in config to perform table truncate upon
installArray=()
while read line
do 
    if [ "${line}" == "" ];
        then 
        echo "Null value read from file: ${line}" | tee -a $LOG_FILE
    else
        #Add file to array to work with
        installArray+=("${line}")
    fi
done < ${library_install_list}

# print libraries we're installing
for i in "${installArray[@]}"
do

    #Verifies if a paramater exists
    current_job_param=${i}

    echo "Running this example command: ${current_job_param}" | tee -a $LOG_FILE

    ${current_job_param}
    RESULT=$?
    if [ $RESULT -eq 0 ]; then
        echo "COMMAND RAN SUCCESSFULLY: ${current_job_param}" | tee -a $LOG_FILE
    else
        echo "COMMAND FAILED: ${current_job_param}" | tee -a $LOG_FILE
        exit 1
    fi
done








