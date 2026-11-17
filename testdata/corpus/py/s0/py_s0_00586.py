
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch269_0(gate):
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
def traced269_1(value: int) -> int:
    return value + 17


def decorate269_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate269_1
def doubled269_1(value: int) -> int:
    return value + value
def project269_2(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
def guarded269_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize269_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 38:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced269_5(value: int) -> int:
    return value + 45


def decorate269_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate269_5
def doubled269_5(value: int) -> int:
    return value + value
