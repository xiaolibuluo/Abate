
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2646_0(values):
    for item in values:
        if item > 65:
            yield item * 48


def collect2646_0(values):
    return list(expand2646_0(values))
def guarded2646_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2646_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 79}
def expand2646_3(values):
    for item in values:
        if item > 86:
            yield item * 20


def collect2646_3(values):
    return list(expand2646_3(values))
def pair2646_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 93}
def pair2646_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 100}
