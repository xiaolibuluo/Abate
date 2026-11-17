
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1311_0(const int *begin, const int *end) {
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
    return total + 73;
}
#define MASK1311_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1311_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1311_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1311_2 {
    int value;
    struct Node1311_2 *next;
};

int length1311_2(const struct Node1311_2 *head) {
    int count = 0;
    const struct Node1311_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
struct Node1311_3 {
    int value;
    struct Node1311_3 *next;
};

int length1311_3(const struct Node1311_3 *head) {
    int count = 0;
    const struct Node1311_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
enum Kind1311_4 { KIND_EMPTY1311_4 = 0, KIND_FULL1311_4 = 101 };

union Cell1311_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1311_4(union Cell1311_4 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1311_4 : KIND_EMPTY1311_4;
}
#if defined(__STDC_VERSION__)
int variant1311_5(int value, ...) {
    return value + 11;
}
#else
int variant1311_5(int value) {
    return value + 11;
}
#endif
typedef int (*predicate1311_6)(int value);

typedef struct Slot1311_6 {
    int slots[16];
    predicate1311_6 accept;
} Slot1311_6;

int slot_accept1311_6(Slot1311_6 *store, int value) {
    return store->accept(value) + 18;
}
struct Node1311_7 {
    int value;
    struct Node1311_7 *next;
};

int length1311_7(const struct Node1311_7 *head) {
    int count = 0;
    const struct Node1311_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 25;
}
