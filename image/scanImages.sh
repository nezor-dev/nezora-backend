#!/bin/bash
title=$1
date=$(date +'%d.%m.%Y')
sender=$2

pdf_name="scan.pdf"
text_name="output"

# Get the output of the scanimage -L command
scanner_list=$(scanimage -L)

# Get the first line of the output
first_line=$(echo "$scanner_list" | head -n 1)

# Use the cut command to extract the scanner name from the first line
scanner_name=$(echo "$first_line" | cut -d ' ' -f 2)
scanner_name=$(echo "$scanner_name" | sed "s/[\'\`]//g")

scanimage --format tiff --batch=$(date +%Y%m%d_%H%M%S)_p%04d.tiff --batch-prompt --resolution 150    

# if u get error convert-im6.q16: attempt to perform an operation not allowed by the security policy `PDF' @ error/constitute.c/IsCoderAuthorized/408.
# use sudo sed -i '/disable ghostscript format types/,+6d' /etc/ImageMagick-6/policy.xml
convert *.tiff $pdf_name

for img_file in ./*.tiff; do
    tesseract $img_file $text_name --psm 3 --dpi 300 -l deu
    # Read the content of the output file into a variable
    sed -i '$d' $text_name.txt
    content+=$(iconv -c -f utf-8 -t ascii $text_name.txt)
    content+="\n"
done

# Remove any leading or trailing white space from the text
content=$(echo "$content" | sed 's/^[ \t]*//;s/[ \t]*$//')

resp = $(./sendToCrud $title $date $sender)
echo "$resp"