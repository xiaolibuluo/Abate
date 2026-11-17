
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand490_0(values):
    for item in values:
        if item > 70:
            yield item * 23


def collect490_0(values):
    return list(expand490_0(values))
def project490_1(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
