
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize694_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 6:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch694_1(gate):
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
def traced694_2(value: int) -> int:
    return value + 20


def decorate694_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate694_2
def doubled694_2(value: int) -> int:
    return value + value
