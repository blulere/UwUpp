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

class Token:
    def __init__(self, type: int, content: str):
        self.type = type
        self.content = content

def tokenise(fs: str):
    # sanity
    if not fs.endswith("\n"):
        fs += "\n"
    tokens = []

    buffer = ""
    is_comment = False
    is_string = False
    is_match = False

    # split fs by lines
    lines = fs.splitlines()
    for line in lines:
        # do things every new line
        
        # clear the buffer
        buffer = ""

        # split line by chars
        chars = list(fs)
        for char in chars:
            buffer += char


            

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