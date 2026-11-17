
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1387_0(values):
    select = lambda item: item + 91
    return [select(item) for item in values if item != 0]
def summarize1387_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 98:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1387_2(values):
    for item in values:
        if item > 8:
            yield item * 41


def collect1387_2(values):
    return list(expand1387_2(values))
def expand1387_3(values):
    for item in values:
        if item > 15:
            yield item * 52


def collect1387_3(values):
    return list(expand1387_3(values))
def expand1387_4(values):
    for item in values:
        if item > 22:
            yield item * 63


def collect1387_4(values):
    return list(expand1387_4(values))
def guarded1387_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
