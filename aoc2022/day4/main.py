
def two(s:list[str]):
    sum = 0

    for l in s:
        p1, p2 = (n for n in l.split(","))
        s1,e1 = (int(n) for n in p1.split("-"))
        s2,e2 = (int(n) for n in p2.split("-"))

        if (s1 <= s2 <= e1) or (s2 <= s1 <= e2) or (s1 <= e2 <= e1) or (s2 <= e1 <= e2):
            sum+=1

    print(sum)



def one(s: list[str]):
    sum = 0
    for l in s:
        p1, p2 = (n for n in l.split(","))
        s1,e1 = (int(n) for n in p1.split("-"))
        s2,e2 = (int(n) for n in p2.split("-"))

        if s1 <= s2 and e2 <= e1:
            sum+=1
        elif s2 <= s1 and e1 <= e2:
            sum += 1

    print(sum)

if __name__ == "__main__":
    s = open("in.txt").read().split("\n")[:-1]

    one(s)
    two(s)
