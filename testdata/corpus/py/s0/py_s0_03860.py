
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1781_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1781_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1781_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 86}
def project1781_3(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
def traced1781_4(value: int) -> int:
    return value + 100


def decorate1781_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1781_4
def doubled1781_4(value: int) -> int:
    return value + value
async def fetch1781_5(gate):
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
def traced1781_6(value: int) -> int:
    return value + 17


def decorate1781_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1781_6
def doubled1781_6(value: int) -> int:
    return value + value
