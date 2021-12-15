from collections import defaultdict

fd = open(0)
p = fd.readline().strip()
fd.readline()
r = dict([x.strip().split(' -> ') for x in fd])


def t(s):
    return [s[i:i+2] for i in range(len(s)-1)]


n = defaultdict(int, dict([[x, 1] for x in t(p)]))

for _ in range(40):
    m = {**n}
    for x in [x for x in n if n[x] > 0 and x in r]:
        n[x] -= m[x]
        for y in t("".join([x[0], r[x], x[1]])):
            n[y] += m[x]

c = defaultdict(int)
for x in n:
    i = n[x]
    c[x[0]] += i

c[p[-1]] += 1

print(max(c.values()) - min(c.values()))
