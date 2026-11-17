
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2092_0(values):
    for item in values:
        if item > 41:
            yield item * 13


def collect2092_0(values):
    return list(expand2092_0(values))
