
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded979_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize979_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 32:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
