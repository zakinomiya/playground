
def one(list: str, n: int):
    
    arr = []
    for i, c in enumerate(list):
        arr.insert(0, c)

        duplicate = False
        for k, a in enumerate(arr):
            for j in range(k+1, len(arr)):
                if (arr[j] == a):
                    duplicate = True
                    break

        if len(arr) == n:
            if duplicate == False:
                print(i+1) 
                return
            else:
                arr.pop()
         
            

if __name__ == "__main__":
    s = open("in").read()

    one(s, 4)
    one(s, 14)
