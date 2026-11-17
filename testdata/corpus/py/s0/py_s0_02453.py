
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1132_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1132_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1132_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 88}
def traced1132_3(value: int) -> int:
    return value + 95


def decorate1132_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1132_3
def doubled1132_3(value: int) -> int:
    return value + value
