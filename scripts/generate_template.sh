#!/bin/bash

url=$1

if [ -z "$url" ]; then
    echo "Error: No URL provided."
    exit 1
fi

if ! [[ $url =~ ^https://leetcode.com/problems/.+ ]]; then
    echo "Error: Invalid URL format."
    exit 1
fi

problem_name=$(echo $url | sed -n 's|.*/problems/\(.*\)/.*|\1|p')

problem_name=${problem_name%/description}

snake_case_filename=$(echo $problem_name | tr '-' '_')

mkdir -p problems/$snake_case_filename
touch problems/$snake_case_filename/main.go

cat << EOS > problems/$snake_case_filename/main.go
package main

import "fmt"

// ${url}

/*
TODO: paste description from leetcode
*/

func dummy() {
    // TODO: implement
}

func main() {
    fmt.Println(dummy()) // Output:
    fmt.Println(dummy()) // Output:
}
EOS
