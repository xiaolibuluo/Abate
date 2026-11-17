
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand446_0(values):
    for item in values:
        if item > 80:
            yield item * 20


def collect446_0(values):
    return list(expand446_0(values))
def expand446_1(values):
    for item in values:
        if item > 87:
            yield item * 31


def collect446_1(values):
    return list(expand446_1(values))
def expand446_2(values):
    for item in values:
        if item > 94:
            yield item * 42


def collect446_2(values):
    return list(expand446_2(values))
