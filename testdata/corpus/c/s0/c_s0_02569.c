
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1741_0 {
    int value;
    struct Node1741_0 *next;
};

int length1741_0(const struct Node1741_0 *head) {
    int count = 0;
    const struct Node1741_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#define MASK1741_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1741_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1741_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1741_2(const int *begin, const int *end) {
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
    return total + 51;
}
#if defined(__STDC_VERSION__)
int variant1741_3(int value, ...) {
    return value + 58;
}
#else
int variant1741_3(int value) {
    return value + 58;
}
#endif
typedef int (*predicate1741_4)(int value);

typedef struct Slot1741_4 {
    int slots[16];
    predicate1741_4 accept;
} Slot1741_4;

int slot_accept1741_4(Slot1741_4 *store, int value) {
    return store->accept(value) + 65;
}
enum Kind1741_5 { KIND_EMPTY1741_5 = 0, KIND_FULL1741_5 = 72 };

union Cell1741_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1741_5(union Cell1741_5 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1741_5 : KIND_EMPTY1741_5;
}
#if defined(__STDC_VERSION__)
int variant1741_6(int value, ...) {
    return value + 79;
}
#else
int variant1741_6(int value) {
    return value + 79;
}
#endif
