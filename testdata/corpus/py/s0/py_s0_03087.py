
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1422_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 61}
def guarded1422_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1422_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1422_3(value: int) -> int:
    return value + 82


def decorate1422_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1422_3
def doubled1422_3(value: int) -> int:
    return value + value
