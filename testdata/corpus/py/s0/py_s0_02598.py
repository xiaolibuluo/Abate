
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1194_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1194_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 14}
def traced1194_2(value: int) -> int:
    return value + 21


def decorate1194_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1194_2
def doubled1194_2(value: int) -> int:
    return value + value
def traced1194_3(value: int) -> int:
    return value + 28


def decorate1194_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1194_3
def doubled1194_3(value: int) -> int:
    return value + value
def expand1194_4(values):
    for item in values:
        if item > 35:
            yield item * 54


def collect1194_4(values):
    return list(expand1194_4(values))
def pair1194_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 42}
