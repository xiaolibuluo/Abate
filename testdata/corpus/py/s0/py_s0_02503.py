
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1158_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1158_1(values):
    for item in values:
        if item > 31:
            yield item * 13


def collect1158_1(values):
    return list(expand1158_1(values))
def expand1158_2(values):
    for item in values:
        if item > 38:
            yield item * 24


def collect1158_2(values):
    return list(expand1158_2(values))
def project1158_3(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
def traced1158_4(value: int) -> int:
    return value + 52


def decorate1158_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1158_4
def doubled1158_4(value: int) -> int:
    return value + value
def project1158_5(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def traced1158_6(value: int) -> int:
    return value + 66


def decorate1158_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1158_6
def doubled1158_6(value: int) -> int:
    return value + value
async def fetch1158_7(gate):
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
def guarded1158_8(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
