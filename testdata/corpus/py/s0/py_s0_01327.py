
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize601_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 58:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced601_1(value: int) -> int:
    return value + 65


def decorate601_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate601_1
def doubled601_1(value: int) -> int:
    return value + value
def summarize601_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 72:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch601_3(gate):
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
