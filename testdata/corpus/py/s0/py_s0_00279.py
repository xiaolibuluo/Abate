
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand125_0(values):
    for item in values:
        if item > 78:
            yield item * 30


def collect125_0(values):
    return list(expand125_0(values))
def pair125_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
def guarded125_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced125_3(value: int) -> int:
    return value + 99


def decorate125_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate125_3
def doubled125_3(value: int) -> int:
    return value + value
def expand125_4(values):
    for item in values:
        if item > 9:
            yield item * 13


def collect125_4(values):
    return list(expand125_4(values))
def project125_5(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def traced125_6(value: int) -> int:
    return value + 23


def decorate125_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate125_6
def doubled125_6(value: int) -> int:
    return value + value
