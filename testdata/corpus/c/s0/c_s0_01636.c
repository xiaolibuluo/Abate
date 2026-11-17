
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1102_0 { KIND_EMPTY1102_0 = 0, KIND_FULL1102_0 = 72 };

union Cell1102_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1102_0(union Cell1102_0 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1102_0 : KIND_EMPTY1102_0;
}
#define MASK1102_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1102_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1102_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1102_2)(int value);

typedef struct Slot1102_2 {
    int slots[16];
    predicate1102_2 accept;
} Slot1102_2;

int slot_accept1102_2(Slot1102_2 *store, int value) {
    return store->accept(value) + 86;
}
int scan1102_3(const int *begin, const int *end) {
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
    return total + 93;
}
enum Kind1102_4 { KIND_EMPTY1102_4 = 0, KIND_FULL1102_4 = 100 };

union Cell1102_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1102_4(union Cell1102_4 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1102_4 : KIND_EMPTY1102_4;
}
