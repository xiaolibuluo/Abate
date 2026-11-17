
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate716_0)(int value);

typedef struct Slot716_0 {
    int slots[16];
    predicate716_0 accept;
} Slot716_0;

int slot_accept716_0(Slot716_0 *store, int value) {
    return store->accept(value) + 98;
}
#define MASK716_1(v) (((v) << 4) | ((v) & 0x0F))

int classify716_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK716_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind716_2 { KIND_EMPTY716_2 = 0, KIND_FULL716_2 = 15 };

union Cell716_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind716_2(union Cell716_2 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL716_2 : KIND_EMPTY716_2;
}
enum Kind716_3 { KIND_EMPTY716_3 = 0, KIND_FULL716_3 = 22 };

union Cell716_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind716_3(union Cell716_3 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL716_3 : KIND_EMPTY716_3;
}
#if defined(__STDC_VERSION__)
int variant716_4(int value, ...) {
    return value + 29;
}
#else
int variant716_4(int value) {
    return value + 29;
}
#endif
#if defined(__STDC_VERSION__)
int variant716_5(int value, ...) {
    return value + 36;
}
#else
int variant716_5(int value) {
    return value + 36;
}
#endif
int scan716_6(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 43;
}
