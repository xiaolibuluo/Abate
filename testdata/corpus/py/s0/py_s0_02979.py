
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1364_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 83}
def expand1364_1(values):
    for item in values:
        if item > 90:
            yield item * 52


def collect1364_1(values):
    return list(expand1364_1(values))
