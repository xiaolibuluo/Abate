
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1905_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1905_1(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
def expand1905_2(values):
    for item in values:
        if item > 49:
            yield item * 68


def collect1905_2(values):
    return list(expand1905_2(values))
