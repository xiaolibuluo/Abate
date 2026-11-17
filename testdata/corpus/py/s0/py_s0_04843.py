
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2240_0(values):
    for item in values:
        if item > 25:
            yield item * 12


def collect2240_0(values):
    return list(expand2240_0(values))
def project2240_1(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
