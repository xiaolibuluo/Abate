
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1486_0(gate):
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
def traced1486_1(value: int) -> int:
    return value + 27


def decorate1486_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1486_1
def doubled1486_1(value: int) -> int:
    return value + value
def traced1486_2(value: int) -> int:
    return value + 34


def decorate1486_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1486_2
def doubled1486_2(value: int) -> int:
    return value + value
def summarize1486_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1486_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 48:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
