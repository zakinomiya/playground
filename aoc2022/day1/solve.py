print(sum(sorted([sum([ int(k) if k != "" else 0 for k in n.split("\n")]) for n in open("in.txt").read().split("\n\n")])[-3:]))
