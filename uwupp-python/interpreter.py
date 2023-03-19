def uwupp_to_python(fs: str):
    pythoncode = ''

    # split fs into lines, then into words.
    lineslist = fs.split('\n')
    
    current_line = 1
    current_word = 1

    for line in lineslist:
        # each new line is automatically not a comment
        # until it comes across a comment.
        is_comment = False
        # ignore most rules as currently parsing a string
        is_string = False
        # reset the current word of the line
        current_word = 1

        wordslist = line.split(' ')
        
        if is_comment == False:
            # this line was never a comment...
            if 'UwU' in line:
                # but now the rest of it is
                is_comment = True
                line = line.replace('UwU', '# UwU')
        
        # errors
        if '*notices i iws 0' in line:
            raise Exception('error: for loops are not supported currently. please convert this to a while loop')
        # warnings
        if 'gweatew twan' in line:
            print(f"warning: this uses issue syntax from the original version of UwU++ (gweatew twan). make sure to fix this to avoid deprecation (line {current_line})")
            line = line.replace('gweatew twan', '>')
        if 'wess twan' in line:
            print(f"warning: this uses issue syntax from the original version of UwU++ (wess twan). make sure to fix this to avoid deprecation (line {current_line})")
            line = line.replace('wess twan', '<')
        if 'eqwall twoo' in line:
            print(f"warning: this uses issue syntax from the original version of UwU++ (eqwall twoo). make sure to fix this to avoid deprecation (line {current_line})")
            line = line.replace('eqwall twoo', '==')
        if 'nuzzels(' in line:
            print(f"warning: this uses issue syntax from the original version of UwU++ (nuzzels). make sure to fix this to avoid deprecation (line {current_line})")
            line = line.replace('nuzzels(', 'print(')

        # TODO: Add checks for strings so that their contents aren't modified (this is KILLING me !!!)

        if is_string == False:
            line = line.replace('iws', '=')
            line = line.replace('OwO *notices', 'while')
            line = line.replace('ewse *notices', 'elif')
            line = line.replace('ewse', 'else:')
            line = line.replace('*notices', 'if')
            line = line.replace('wess than', '<')
            line = line.replace('gweatew than', '>')
            line = line.replace('not eqwaws', '!=')
            line = line.replace('*', ':')
            line = line.replace('twimes', '*')
            line = line.replace('moduwo', '%')
            line = line.replace('eqwaw', '==')
            line = line.replace('diwide', '/')
            line = line.replace('nuzzles(', 'print(')
            
            """
            Python doesn't support fixed array allocation.
            This is a workaround.
            So far it works only with Strings, Ints
            and Nones.
            """
            if 'awway<' in line:
                if 'stwing' in line:
                    line = line.replace('awway<', '[""]*')
                elif 'int' in line:
                    line = line.replace('awway<', '[0]*')
                else:
                    line = line.replace('awway<', '[None]*')
                line = line.replace('>', '')
            line = line.replace('stwing', 'str')
            line = line.replace('stawp', '')
            line = line.replace('pwus', '+')

        pythoncode += line

        # increment the word counter
        # and add a space because bug
        # current_word += 1

        # incremenet the line counter
        # and add a newline because bug
        pythoncode += '\n'
        current_line += 1
    
    return pythoncode

def python_to_uwupp(fs):
    # coming soon
    pass

def interpreter(fs):
    """
    Interpret the code directly without having to make
    a new Python file
    """
    # coming soon
    pass