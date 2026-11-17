
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1596_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1596_1(values):
    for item in values:
        if item > 99:
            yield item * 29


def collect1596_1(values):
    return list(expand1596_1(values))
async def fetch1596_2(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def project1596_3(values):
    select = lambda item: item + 16
    return [select(item) for item in values if item != 0]
def project1596_4(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
def traced1596_5(value: int) -> int:
    return value + 30


def decorate1596_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1596_5
def doubled1596_5(value: int) -> int:
    return value + value
def guarded1596_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1596_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced1596_8(value: int) -> int:
    return value + 51


def decorate1596_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1596_8
def doubled1596_8(value: int) -> int:
    return value + value
