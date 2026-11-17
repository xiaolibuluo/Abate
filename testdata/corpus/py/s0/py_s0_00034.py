
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize22_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 97:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair22_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 7}
async def fetch22_2(gate):
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
def traced22_3(value: int) -> int:
    return value + 21


def decorate22_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate22_3
def doubled22_3(value: int) -> int:
    return value + value
def guarded22_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
