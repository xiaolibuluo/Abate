
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project638_0(values):
    select = lambda item: item + 54
    return [select(item) for item in values if item != 0]
def guarded638_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand638_2(values):
    for item in values:
        if item > 68:
            yield item * 44


def collect638_2(values):
    return list(expand638_2(values))
