
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1179_0(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def pair1179_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def project1179_2(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
def expand1179_3(values):
    for item in values:
        if item > 27:
            yield item * 60


def collect1179_3(values):
    return list(expand1179_3(values))
