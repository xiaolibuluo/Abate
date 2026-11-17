
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1557_0(values):
    for item in values:
        if item > 70:
            yield item * 50


def collect1557_0(values):
    return list(expand1557_0(values))
