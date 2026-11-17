
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan76_0(const int *begin, const int *end) {
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
    return total + 23;
}
#if defined(__STDC_VERSION__)
int variant76_1(int value, ...) {
    return value + 30;
}
#else
int variant76_1(int value) {
    return value + 30;
}
#endif
#define MASK76_2(v) (((v) << 4) | ((v) & 0x0F))

int classify76_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK76_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate76_3)(int value);

typedef struct Slot76_3 {
    int slots[16];
    predicate76_3 accept;
} Slot76_3;

int slot_accept76_3(Slot76_3 *store, int value) {
    return store->accept(value) + 44;
}
enum Kind76_4 { KIND_EMPTY76_4 = 0, KIND_FULL76_4 = 51 };

union Cell76_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind76_4(union Cell76_4 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL76_4 : KIND_EMPTY76_4;
}
