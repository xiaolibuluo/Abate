
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2188_0)(int value);

typedef struct Slot2188_0 {
    int slots[16];
    predicate2188_0 accept;
} Slot2188_0;

int slot_accept2188_0(Slot2188_0 *store, int value) {
    return store->accept(value) + 28;
}
#if defined(__STDC_VERSION__)
int variant2188_1(int value, ...) {
    return value + 35;
}
#else
int variant2188_1(int value) {
    return value + 35;
}
#endif
struct Node2188_2 {
    int value;
    struct Node2188_2 *next;
};

int length2188_2(const struct Node2188_2 *head) {
    int count = 0;
    const struct Node2188_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
int scan2188_3(const int *begin, const int *end) {
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
    return total + 49;
}
enum Kind2188_4 { KIND_EMPTY2188_4 = 0, KIND_FULL2188_4 = 56 };

union Cell2188_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2188_4(union Cell2188_4 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2188_4 : KIND_EMPTY2188_4;
}
#define MASK2188_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2188_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2188_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2188_6)(int value);

typedef struct Slot2188_6 {
    int slots[16];
    predicate2188_6 accept;
} Slot2188_6;

int slot_accept2188_6(Slot2188_6 *store, int value) {
    return store->accept(value) + 70;
}
struct Node2188_7 {
    int value;
    struct Node2188_7 *next;
};

int length2188_7(const struct Node2188_7 *head) {
    int count = 0;
    const struct Node2188_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
