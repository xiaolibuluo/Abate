
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1925_0(const int *begin, const int *end) {
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
    return total + 101;
}
enum Kind1925_1 { KIND_EMPTY1925_1 = 0, KIND_FULL1925_1 = 11 };

union Cell1925_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1925_1(union Cell1925_1 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1925_1 : KIND_EMPTY1925_1;
}
typedef int (*predicate1925_2)(int value);

typedef struct Slot1925_2 {
    int slots[16];
    predicate1925_2 accept;
} Slot1925_2;

int slot_accept1925_2(Slot1925_2 *store, int value) {
    return store->accept(value) + 18;
}
int scan1925_3(const int *begin, const int *end) {
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
    return total + 25;
}
#define MASK1925_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1925_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1925_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
