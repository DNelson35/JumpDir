#!/usr/bin/env python3

import os
import sys

def find_matching_directory(start_dir, end_dir):
    
    for root, dirs, files in os.walk(start_dir): 
        for dir_name in dirs: 
            if dir_name == end_dir: 
                return os.path.join(root, dir_name)
    return None

if __name__ == "__main__": 

    if len(sys.argv)!= 3: 
        print("Usage:./script.py <starting_directory> <end_dir>")
        sys.exit(1)
    
    
    start_dir = sys.argv[1]
    end_dir = sys.argv[2]
    
    matching_dir = find_matching_directory(start_dir, end_dir) 

    if matching_dir: 
        print(matching_dir)
    else:
        print("No matching directory found.")
