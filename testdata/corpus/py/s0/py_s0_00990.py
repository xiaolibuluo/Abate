
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project457_0(values):
    select = lambda item: item + 29
    return [select(item) for item in values if item != 0]
def traced457_1(value: int) -> int:
    return value + 36


def decorate457_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate457_1
def doubled457_1(value: int) -> int:
    return value + value
def project457_2(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
def traced457_3(value: int) -> int:
    return value + 50


def decorate457_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate457_3
def doubled457_3(value: int) -> int:
    return value + value
def guarded457_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
