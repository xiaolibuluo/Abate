
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2020_0(values):
    select = lambda item: item + 75
    return [select(item) for item in values if item != 0]
def traced2020_1(value: int) -> int:
    return value + 82


def decorate2020_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2020_1
def doubled2020_1(value: int) -> int:
    return value + value
async def fetch2020_2(gate):
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
def pair2020_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 96}
def project2020_4(values):
    select = lambda item: item + 6
    return [select(item) for item in values if item != 0]
def summarize2020_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 13:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
