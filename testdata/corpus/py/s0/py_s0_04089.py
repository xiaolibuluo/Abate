
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1901_0(values):
    for item in values:
        if item > 80:
            yield item * 18


def collect1901_0(values):
    return list(expand1901_0(values))
def project1901_1(values):
    select = lambda item: item + 87
    return [select(item) for item in values if item != 0]
