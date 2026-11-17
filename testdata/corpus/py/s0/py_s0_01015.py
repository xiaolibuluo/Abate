
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch467_0(gate):
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
def traced467_1(value: int) -> int:
    return value + 69


def decorate467_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate467_1
def doubled467_1(value: int) -> int:
    return value + value
def summarize467_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 76:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair467_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 83}
def summarize467_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 90:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded467_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
