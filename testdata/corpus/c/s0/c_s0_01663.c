
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1117_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1117_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1117_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1117_1)(int value);

typedef struct Slot1117_1 {
    int slots[16];
    predicate1117_1 accept;
} Slot1117_1;

int slot_accept1117_1(Slot1117_1 *store, int value) {
    return store->accept(value) + 80;
}
#if defined(__STDC_VERSION__)
int variant1117_2(int value, ...) {
    return value + 87;
}
#else
int variant1117_2(int value) {
    return value + 87;
}
#endif
int scan1117_3(const int *begin, const int *end) {
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
    return total + 94;
}
enum Kind1117_4 { KIND_EMPTY1117_4 = 0, KIND_FULL1117_4 = 101 };

union Cell1117_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1117_4(union Cell1117_4 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1117_4 : KIND_EMPTY1117_4;
}
struct Node1117_5 {
    int value;
    struct Node1117_5 *next;
};

int length1117_5(const struct Node1117_5 *head) {
    int count = 0;
    const struct Node1117_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node1117_6 {
    int value;
    struct Node1117_6 *next;
};

int length1117_6(const struct Node1117_6 *head) {
    int count = 0;
    const struct Node1117_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
struct Node1117_7 {
    int value;
    struct Node1117_7 *next;
};

int length1117_7(const struct Node1117_7 *head) {
    int count = 0;
    const struct Node1117_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
