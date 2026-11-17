
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project831_0(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
def expand831_1(values):
    for item in values:
        if item > 48:
            yield item * 42


def collect831_1(values):
    return list(expand831_1(values))
def expand831_2(values):
    for item in values:
        if item > 55:
            yield item * 53


def collect831_2(values):
    return list(expand831_2(values))
