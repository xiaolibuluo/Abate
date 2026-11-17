
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded926_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced926_1(value: int) -> int:
    return value + 22


def decorate926_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate926_1
def doubled926_1(value: int) -> int:
    return value + value
