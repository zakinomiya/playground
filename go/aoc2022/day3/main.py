

def get_value(c: str) -> int:
    if ord(c) < ord('a'):
        return ord(c) - 38
    else:
        return ord(c) - 96

def one(strings: list[str]):
    sum = 0
    for ss in strings:
        m = len(ss) // 2
        l = {}
        f, s = ss[:m], ss[m:]

        for c in f:
            l[c] = 1

        for c in s:
            if c in l:
                sum += get_value(c)
                break

    print(sum)

def two(strings: list[str]):
    sum = 0
    for i in range(0, len(strings), 3):
        l = {}
        for s in strings[i]:
            l[s] = 1

        for s in strings[i+1]:
            if s in l and l[s] == 1:
                l[s] += 1

        for s in strings[i+2]:
            if s in l and l[s] == 2:
                sum += get_value(s)
                break
    print(sum)

if __name__ == "__main__":
    strings = open("in.txt").read().split("\n")[:-1]

    one(strings)
    two(strings)
