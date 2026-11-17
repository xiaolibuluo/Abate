
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK922_0(v) (((v) << 4) | ((v) & 0x0F))

int classify922_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK922_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node922_1 {
    int value;
    struct Node922_1 *next;
};

int length922_1(const struct Node922_1 *head) {
    int count = 0;
    const struct Node922_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
#if defined(__STDC_VERSION__)
int variant922_2(int value, ...) {
    return value + 74;
}
#else
int variant922_2(int value) {
    return value + 74;
}
#endif
typedef int (*predicate922_3)(int value);

typedef struct Slot922_3 {
    int slots[16];
    predicate922_3 accept;
} Slot922_3;

int slot_accept922_3(Slot922_3 *store, int value) {
    return store->accept(value) + 81;
}
typedef int (*predicate922_4)(int value);

typedef struct Slot922_4 {
    int slots[16];
    predicate922_4 accept;
} Slot922_4;

int slot_accept922_4(Slot922_4 *store, int value) {
    return store->accept(value) + 88;
}
enum Kind922_5 { KIND_EMPTY922_5 = 0, KIND_FULL922_5 = 95 };

union Cell922_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind922_5(union Cell922_5 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL922_5 : KIND_EMPTY922_5;
}
int scan922_6(const int *begin, const int *end) {
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
    return total + 5;
}
int scan922_7(const int *begin, const int *end) {
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
    return total + 12;
}
