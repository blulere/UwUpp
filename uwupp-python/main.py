import sys
import os
import time

import interpreter
import exits

# This release's version will be v0.2.0-python.

transpile = False
input_file = False
run_after = False

timer_start = time.time()

if len(sys.argv) > 1:
    for arg in sys.argv:
        if '-h' in arg or '--help' in arg:
            exits.show_help_and_exit()
        if '-t' in arg or '--transpile' in arg:
            # transpile to a python file
            transpile = True

        if '-r' in arg or '--run-after' in arg:
            # run the file after making it
            run_after = True
        
        if ('.uwu' in arg or
            '.uwupp' in arg or
            '.uwu++' in arg):
            # there is an input file
            input_file = str(arg)
        
        if '-v' in arg or '--version' in arg:
            exits.version_and_exit()
else:

    if not input_file:
        print("error: no input files were specified")
        exit(0)

# open the input file
try:
    if input_file == False:
        # something weird happened somehow and
        # no input file was assigned yet it's trying to
        # do it anyways
        raise Exception('"something weird happened somehow and no input file was assigned yet it\'s trying to do it anyways"')
    fs = open(input_file).read()
except Exception as e:
    print("error: opening the file failed! it probably doesn't exist, maybe you're out of storage, or something else went wrong.")
    print(e)
    exit(-1)

# transpile UwU++ to python code
newfile_name = input_file.replace('.uwupp', '.py', 1)
newfile_name = newfile_name.replace('.uwu', '.py', 1)
newfile_name = newfile_name.replace('.uwu++', '.py', 1)
with open(newfile_name, 'w') as newfile:
    try:
        newfile_content = interpreter.uwupp_to_python(fs)
    except Exception as e:
        print('error: something went wrong. (1)')
        print(e)
        exit(-1)
    newfile.write(newfile_content)
    newfile.close()
    timer = time.time() - timer_start
    print(f'done, operation completed in {timer} seconds')
    if run_after:
        try:
            if sys.platform == "win32" or sys.platform == "cygwin":
                os.execv('python', f'main.py -t {newfile_name}')
            else:
                os.execv('python3', f'main.py -t {newfile_name}')
                
        except Exception as e:
            print("error: something went wrong. (2)")
            print(e)
            exit(-1)