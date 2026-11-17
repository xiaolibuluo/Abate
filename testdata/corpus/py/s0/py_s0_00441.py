
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand202_0(values):
    for item in values:
        if item > 12:
            yield item * 20


def collect202_0(values):
    return list(expand202_0(values))
def project202_1(values):
    select = lambda item: item + 19
    return [select(item) for item in values if item != 0]
def guarded202_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
