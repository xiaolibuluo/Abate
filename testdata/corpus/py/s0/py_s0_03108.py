
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1433_0(value: int) -> int:
    return value + 10


def decorate1433_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1433_0
def doubled1433_0(value: int) -> int:
    return value + value
def project1433_1(values):
    select = lambda item: item + 17
    return [select(item) for item in values if item != 0]
def summarize1433_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 24:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1433_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 31}
async def fetch1433_4(gate):
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
