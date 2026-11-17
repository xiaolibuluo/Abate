
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2313_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair2313_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 11}
def pair2313_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 18}
def traced2313_3(value: int) -> int:
    return value + 25


def decorate2313_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2313_3
def doubled2313_3(value: int) -> int:
    return value + value
def project2313_4(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
def expand2313_5(values):
    for item in values:
        if item > 39:
            yield item * 29


def collect2313_5(values):
    return list(expand2313_5(values))
def traced2313_6(value: int) -> int:
    return value + 46


def decorate2313_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2313_6
def doubled2313_6(value: int) -> int:
    return value + value
