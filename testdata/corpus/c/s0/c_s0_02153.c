
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1449_0 { KIND_EMPTY1449_0 = 0, KIND_FULL1449_0 = 24 };

union Cell1449_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1449_0(union Cell1449_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1449_0 : KIND_EMPTY1449_0;
}
struct Node1449_1 {
    int value;
    struct Node1449_1 *next;
};

int length1449_1(const struct Node1449_1 *head) {
    int count = 0;
    const struct Node1449_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
#define MASK1449_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1449_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1449_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1449_3(int value, ...) {
    return value + 45;
}
#else
int variant1449_3(int value) {
    return value + 45;
}
#endif
int scan1449_4(const int *begin, const int *end) {
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
    return total + 52;
}
typedef int (*predicate1449_5)(int value);

typedef struct Slot1449_5 {
    int slots[16];
    predicate1449_5 accept;
} Slot1449_5;

int slot_accept1449_5(Slot1449_5 *store, int value) {
    return store->accept(value) + 59;
}
#define MASK1449_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1449_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1449_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
