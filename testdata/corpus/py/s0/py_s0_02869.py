
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1316_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 41}
def expand1316_1(values):
    for item in values:
        if item > 48:
            yield item * 21


def collect1316_1(values):
    return list(expand1316_1(values))
