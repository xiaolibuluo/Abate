
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded880_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project880_1(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def traced880_2(value: int) -> int:
    return value + 13


def decorate880_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate880_2
def doubled880_2(value: int) -> int:
    return value + value
