import os
from os.path import isdir

files = os.listdir()

for file in files:
    if isdir(file):
        continue
    if file == "sort.py":
        continue
    # move that file to folder
    path = str(file).split(".")[0]
    if not os.path.exists(path):
        os.makedirs(path)
    
    os.rename(file, path + "\\" + file)
