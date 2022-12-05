import os


a = [
    ['V', 'R', 'H', 'B', 'G', "D" ,"W"],
    ["F", "R", "C", "G", "N", "J"],
    ["J", "N", "D", "H", "F", "S", "L"],
    ["V", "S", "D", "J"], 
    ["V", "N", "W", "Q", "R", "D", "H", "S"],
    ["M", "C", "H", "G", "P"],
    ["C", "H", "Z", "L", "G", "B", "J", "F"],
    ["R", "J", "S"],
    ["M", "V", "N", "B", "R", "S", "G", "L"]
]

def show():
    res = []
    for s in a:
        if len(s) > 0:
            res.append(s[0])

    print(a)
    print("".join(res))

def two(s: list[str]):
    for ss in s:
        l = ss.split(" ")
        target, fr, to = int(l[1]), int(l[3])-1, int(l[5])-1
        
        m = target
        if len(a[fr]) <= m:
            m = len(a[fr])

        a[to] = ["-"] * m + a[to]
        a[to][0:m] = a[fr][0:m]
        a[fr] = a[fr][m:]
        # print(a[to])
        # print(a[fr])
        # exit(0)

    show()
    res = 0

    print(res)

def one(s: list[str]):
    for ss in s:
        l = ss.split(" ")
        target, fr, to = int(l[1]), int(l[3])-1, int(l[5])-1

        for i in range(target):
            if len(a[fr]) > 0:
                a[to] = ["-"] + a[to]
                a[to][0] = a[fr][0]
                a[fr] = a[fr][1:]

    show()


if __name__ == "__main__":
    s = open("in.txt").read().split("\n")[:-1]
    # one(s)
    two(s)
