
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1940_0(value: int) -> int:
    return value + 5


def decorate1940_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1940_0
def doubled1940_0(value: int) -> int:
    return value + value
def project1940_1(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def guarded1940_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1940_3(values):
    for item in values:
        if item > 26:
            yield item * 19


def collect1940_3(values):
    return list(expand1940_3(values))
def traced1940_4(value: int) -> int:
    return value + 33


def decorate1940_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1940_4
def doubled1940_4(value: int) -> int:
    return value + value
def pair1940_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 40}
def traced1940_6(value: int) -> int:
    return value + 47


def decorate1940_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1940_6
def doubled1940_6(value: int) -> int:
    return value + value
def traced1940_7(value: int) -> int:
    return value + 54


def decorate1940_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1940_7
def doubled1940_7(value: int) -> int:
    return value + value
def traced1940_8(value: int) -> int:
    return value + 61


def decorate1940_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1940_8
def doubled1940_8(value: int) -> int:
    return value + value
