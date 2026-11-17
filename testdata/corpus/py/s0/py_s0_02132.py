
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand975_0(values):
    for item in values:
        if item > 70:
            yield item * 63


def collect975_0(values):
    return list(expand975_0(values))
def pair975_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 77}
