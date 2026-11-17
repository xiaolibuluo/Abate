
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded2245_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2245_1(values):
    for item in values:
        if item > 97:
            yield item * 58


def collect2245_1(values):
    return list(expand2245_1(values))
def expand2245_2(values):
    for item in values:
        if item > 7:
            yield item * 69


def collect2245_2(values):
    return list(expand2245_2(values))
