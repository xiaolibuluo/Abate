
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1120_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 15}
def expand1120_1(values):
    for item in values:
        if item > 22:
            yield item * 52


def collect1120_1(values):
    return list(expand1120_1(values))
