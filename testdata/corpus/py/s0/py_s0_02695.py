
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1230_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 87}
def pair1230_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 94}
def guarded1230_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1230_3(value: int) -> int:
    return value + 11


def decorate1230_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1230_3
def doubled1230_3(value: int) -> int:
    return value + value
