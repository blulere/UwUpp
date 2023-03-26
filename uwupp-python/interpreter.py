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
    word = ""
    is_comment = False
    is_string = False
    is_match = False

    # split fs by words
    words = fs.split(" ")

    for word in words:
        word += " "
        
        if buffer.endswith("\n"):
            buffer = ""
            is_comment = False
            is_match = False

        # check for comments
        if buffer.startswith("UwU"):
            # beginning of comment definitely
            is_comment = True

        # check for strings
        if buffer.startswith('"') and not buffer.endswith('"'):
            # beginning of string definitely
            is_string = True

        if buffer.endswith('"'):
            # end of string definitely
            is_string = False

        if not is_string and not is_comment:
            if "iws" in buffer:
                tokens.append(Token(ASSIGNMENT, "iws"))
                is_match = True
            elif "OwO *notices" in buffer:
                tokens.append(Token(WHILE, "OwO *notices"))
                is_match = True

        print("Is match: " + str(is_match))
        print("Is string: " + str(is_string))
        print("Is comment: " + str(is_comment))
        print()

        if is_match:
            buffer = ""
            is_match = False
        else:
            buffer += word
            pass

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