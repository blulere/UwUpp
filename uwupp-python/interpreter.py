class Buffer:
    def __init__(self, s, n):
        self.s: str = s
        # next word
        self.n: str = n
    
    def replace(self, rep, _with):
        self.s.replace(rep, _with, -1)


# Token types since they're used by everything anyways,
# no point in putting them away from global scope
COMMENT = 0
VARIABLE = 1
ASSIGNMENT = 2
NUMBER = 3
WHILE = 4
GREATERTHAN = 5
# * at end of loop
STARTLOOP = 6
ADD = 7
# stawp
ENDLOOP = 8
FUNCTIONCALL = 9
OPENPAREN = 10
CLOSEPAREN = 11
EOF = 12
FUNCTIONARG = 13
UNKNOWN = 14
STRING = 15
NEWLINE = 16

class Token:
    def __init__(self, type: int, content: str):
        self.type = type
        self.content = content

def tokenise(fs: str):
    # add a newline at the end of the file for parsing sanity
    if not fs.endswith("\n"):
        fs += "\n"
    tokens = []

    buffer = ""
    is_comment = False
    is_string = False
    in_loop = 0

    # split fs by lines
    lines = fs.splitlines()
    for line in lines:
        # do things every new line
        
        # clear the buffer
        buffer = ""
        
        is_comment = False
        # replace only the first instance in the line,
        # the rest are just comments
        if "UwU " in line:
            line = line.replace("UwU ", "# UwU ", 1)
            is_comment = True

        # split line by words
        words = line.split(" ")
        for word in words:
            # add the space back
            buffer += (word + " ")
            
            # check keywords, if match, add keywords to tokens
            if "iws" in buffer:
                buffer = buffer.replace("iws", "=")
                tokens.append(Token(ASSIGNMENT, buffer))
                buffer = ""
            elif "OwO *notices " in buffer:
                # in_loop += 1
                buffer = buffer.replace("OwO *notices ", "while")
                tokens.append(Token(WHILE, buffer))
                buffer = ""
            elif "gweatew twan" in buffer:
                # this is old syntax
                buffer = buffer.replace("gweatew twan", ">")
                tokens.append(Token(COMMENT, "# Old syntax detected by tokeniser !!!\n"))
                tokens.append(Token(GREATERTHAN, buffer))
                buffer = ""
            elif "*" in buffer:
                in_loop -= 1
                tokens.append(Token(STARTLOOP, ":"))
                buffer = ""
            elif "pwus" in buffer:
                tokens.append(Token(ADD, "+"))
                buffer = ""
            elif "stawp" in buffer:
                tokens.append(Token(ENDLOOP, ""))
                buffer = ""
            elif "nuzzels" in buffer:
                # this is old syntax
                tokens.append(Token(FUNCTIONCALL, "print"))
                buffer = ""
            elif "nuzzles" in buffer:
                tokens.append(Token(FUNCTIONCALL, "print"))
                buffer = ""
            else:
                tokens.append(Token(UNKNOWN, buffer))
                buffer = ""

        tokens.append(Token(NEWLINE, "\n"))

    return tokens


def parse(tokens: list, mode: int = 1):
    final = ""
    
    if mode == 0:
        # silently ignore
        return
    if mode == 1:
        # parse into Python code.
        for token in tokens:
            final += token.content
    
    return final

def uwupp_to_python(fs: str):
    NONE = 0
    PYTHON = 1
    tokens = tokenise(fs)
    pythoncode = str(parse(tokens, PYTHON))

    return pythoncode