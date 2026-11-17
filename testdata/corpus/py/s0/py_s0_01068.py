
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded494_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded494_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced494_2(value: int) -> int:
    return value + 39


def decorate494_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate494_2
def doubled494_2(value: int) -> int:
    return value + value
