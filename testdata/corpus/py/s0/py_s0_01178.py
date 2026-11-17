
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand543_0(values):
    for item in values:
        if item > 80:
            yield item * 28


def collect543_0(values):
    return list(expand543_0(values))
def expand543_1(values):
    for item in values:
        if item > 87:
            yield item * 39


def collect543_1(values):
    return list(expand543_1(values))
def expand543_2(values):
    for item in values:
        if item > 94:
            yield item * 50


def collect543_2(values):
    return list(expand543_2(values))
