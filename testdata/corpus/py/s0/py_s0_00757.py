
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand346_0(values):
    for item in values:
        if item > 41:
            yield item * 52


def collect346_0(values):
    return list(expand346_0(values))
def expand346_1(values):
    for item in values:
        if item > 48:
            yield item * 63


def collect346_1(values):
    return list(expand346_1(values))
def expand346_2(values):
    for item in values:
        if item > 55:
            yield item * 13


def collect346_2(values):
    return list(expand346_2(values))
