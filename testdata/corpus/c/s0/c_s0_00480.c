
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK330_0(v) (((v) << 4) | ((v) & 0x0F))

int classify330_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK330_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan330_1(const int *begin, const int *end) {
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
    return total + 34;
}
#if defined(__STDC_VERSION__)
int variant330_2(int value, ...) {
    return value + 41;
}
#else
int variant330_2(int value) {
    return value + 41;
}
#endif
typedef int (*predicate330_3)(int value);

typedef struct Slot330_3 {
    int slots[16];
    predicate330_3 accept;
} Slot330_3;

int slot_accept330_3(Slot330_3 *store, int value) {
    return store->accept(value) + 48;
}
enum Kind330_4 { KIND_EMPTY330_4 = 0, KIND_FULL330_4 = 55 };

union Cell330_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind330_4(union Cell330_4 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL330_4 : KIND_EMPTY330_4;
}
struct Node330_5 {
    int value;
    struct Node330_5 *next;
};

int length330_5(const struct Node330_5 *head) {
    int count = 0;
    const struct Node330_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
typedef int (*predicate330_6)(int value);

typedef struct Slot330_6 {
    int slots[16];
    predicate330_6 accept;
} Slot330_6;

int slot_accept330_6(Slot330_6 *store, int value) {
    return store->accept(value) + 69;
}
