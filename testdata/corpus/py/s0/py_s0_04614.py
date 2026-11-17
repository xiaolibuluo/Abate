
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2142_0(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def expand2142_1(values):
    for item in values:
        if item > 19:
            yield item * 69


def collect2142_1(values):
    return list(expand2142_1(values))
