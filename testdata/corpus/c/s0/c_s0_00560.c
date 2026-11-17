
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant383_0(int value, ...) {
    return value + 37;
}
#else
int variant383_0(int value) {
    return value + 37;
}
#endif
typedef int (*predicate383_1)(int value);

typedef struct Slot383_1 {
    int slots[16];
    predicate383_1 accept;
} Slot383_1;

int slot_accept383_1(Slot383_1 *store, int value) {
    return store->accept(value) + 44;
}
enum Kind383_2 { KIND_EMPTY383_2 = 0, KIND_FULL383_2 = 51 };

union Cell383_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind383_2(union Cell383_2 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL383_2 : KIND_EMPTY383_2;
}
int scan383_3(const int *begin, const int *end) {
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
    return total + 58;
}
#define MASK383_4(v) (((v) << 4) | ((v) & 0x0F))

int classify383_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK383_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK383_5(v) (((v) << 4) | ((v) & 0x0F))

int classify383_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK383_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate383_6)(int value);

typedef struct Slot383_6 {
    int slots[16];
    predicate383_6 accept;
} Slot383_6;

int slot_accept383_6(Slot383_6 *store, int value) {
    return store->accept(value) + 79;
}
struct Node383_7 {
    int value;
    struct Node383_7 *next;
};

int length383_7(const struct Node383_7 *head) {
    int count = 0;
    const struct Node383_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
