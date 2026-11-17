
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2099_0(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def expand2099_1(values):
    for item in values:
        if item > 42:
            yield item * 12


def collect2099_1(values):
    return list(expand2099_1(values))
def project2099_2(values):
    select = lambda item: item + 49
    return [select(item) for item in values if item != 0]
def expand2099_3(values):
    for item in values:
        if item > 56:
            yield item * 34


def collect2099_3(values):
    return list(expand2099_3(values))
