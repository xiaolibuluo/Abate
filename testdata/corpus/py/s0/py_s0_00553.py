
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded253_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded253_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced253_2(value: int) -> int:
    return value + 10


def decorate253_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate253_2
def doubled253_2(value: int) -> int:
    return value + value
