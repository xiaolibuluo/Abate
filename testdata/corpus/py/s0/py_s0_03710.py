
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1710_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 22}
def pair1710_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 29}
def expand1710_2(values):
    for item in values:
        if item > 36:
            yield item * 45


def collect1710_2(values):
    return list(expand1710_2(values))
