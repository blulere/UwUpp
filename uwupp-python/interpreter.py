# define a mapping of UwU keywords to Python keywords
keywords = {
    'iws': '=',
    'OwO *notices': 'while',
    '*': ':',
    'stawp': 'pass',
    'nuzzels': 'print',
    'pwus': '+',
    'gweatew than': '>',
}

# this should cause a warning when matched, or not
warn_on_legacy = True
legacykeywords = {
    'nuzzels': 'print',
    'gweatew': '>',
    'twan': '',
}

def uwupp_to_python(source: str):
    if not source.endswith('\n'):
        source += '\n'

    pythoncode = ""
    buffer = ""

    lines = source.split('\n')
    for line in lines:
        # re-add a newline at the end
        line += '\n'

        # Replace first instance of UwU in a line with # UwU,
        # this will indicate the start of a single-line comment
        # There is no need to implement termination as
        # this is not specified in language documentation
        line = line.replace("UwU", "# UwU", 1)

        words = line.split(' ')
        for word in words:
            if not word.endswith('\n'):
                word += ' '

            if buffer in keywords:
                pythoncode += keywords[buffer]
                buffer = ""
            elif buffer in legacykeywords:
                pythoncode += legacykeywords[buffer]
                buffer = ""
            else:
                pythoncode += buffer
                buffer = word

        # add any remaining buffer to the python code

    return pythoncode
