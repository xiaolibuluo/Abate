
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair987_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
def traced987_1(value: int) -> int:
    return value + 39


def decorate987_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate987_1
def doubled987_1(value: int) -> int:
    return value + value
def guarded987_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded987_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
