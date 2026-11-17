
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand683_0(values):
    for item in values:
        if item > 57:
            yield item * 32


def collect683_0(values):
    return list(expand683_0(values))
def project683_1(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
