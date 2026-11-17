
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1359_0(int value, ...) {
    return value + 18;
}
#else
int variant1359_0(int value) {
    return value + 18;
}
#endif
#if defined(__STDC_VERSION__)
int variant1359_1(int value, ...) {
    return value + 25;
}
#else
int variant1359_1(int value) {
    return value + 25;
}
#endif
#define MASK1359_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1359_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1359_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1359_3)(int value);

typedef struct Slot1359_3 {
    int slots[16];
    predicate1359_3 accept;
} Slot1359_3;

int slot_accept1359_3(Slot1359_3 *store, int value) {
    return store->accept(value) + 39;
}
#define MASK1359_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1359_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1359_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1359_5 { KIND_EMPTY1359_5 = 0, KIND_FULL1359_5 = 53 };

union Cell1359_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1359_5(union Cell1359_5 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL1359_5 : KIND_EMPTY1359_5;
}
int scan1359_6(const int *begin, const int *end) {
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
    return total + 60;
}
#define MASK1359_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1359_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1359_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
