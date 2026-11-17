
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1463_0(value: int) -> int:
    return value + 12


def decorate1463_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1463_0
def doubled1463_0(value: int) -> int:
    return value + value
def guarded1463_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
