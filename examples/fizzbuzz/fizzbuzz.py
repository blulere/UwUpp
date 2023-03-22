# UwU Fizzbuzz example, transpiled to Python       ?
# UwU This is also a good demonstration of modulus ?
# UwU - blulere                                    ?

# UwU Change this to change when the numbers should stop.
stop_counting: int = 100

i = 0
while i < stop_counting:
    if i % 3 == 0 and i % 5 == 0:
        print("FwizzBwuzz")
    elif i % 3 == 0:
        print("Fwizz")
    elif i % 5 == 0:
        print("Bwuzz")
    else:
        print(i)

    i = i + 1