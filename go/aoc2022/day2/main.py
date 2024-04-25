
rules = {
    'A' : 1, # r
    'B' : 2, # p
    'C' : 3, # s
    'X' : 1, # r
    'Y' : 2, # p
    'Z' : 3, # s
}


def one(arr: list[list[str]]):
    score = 0
    for a in arr:
        if a[0] == '' or a[1] == '':
            continue
        o, m = int(rules[a[0]]), rules[a[1]]
        if m - o == 0:
            score += 3 + m
        elif m-o == 1:
            score += 6 + m
        elif m-o == 2:
            score += 0 + m
        elif m-o == -1:
            score += 0 + m
        elif m-o == -2:
            score += 6 + m
            
    print(score)

def two(arr: list[list[str]]):
    score = 0
    for a in arr:
        if a[0] == '' or a[1] == '':
            continue
        o = a[0]
        if a[1] == 'X':
            score += 0
            if o == 'A':
                score += 3
            elif o == 'B':
                score += 1
            elif o == 'C':
                score += 2
        elif a[1] == 'Y':
            score += 3 + rules[o]
        elif a[1] == 'Z':
            score += 6 
            if o == 'A':
                score += 2
            elif o == 'B':
                score += 3
            elif o == 'C':
                score += 1
    print(score)

if __name__ == "__main__":
    arr = [n.split(" ") for n in open("in.txt").read().split("\n")]
    one(arr)
    two(arr)

