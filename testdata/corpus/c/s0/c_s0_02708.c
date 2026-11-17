
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1841_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1841_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1841_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1841_1 { KIND_EMPTY1841_1 = 0, KIND_FULL1841_1 = 83 };

union Cell1841_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1841_1(union Cell1841_1 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1841_1 : KIND_EMPTY1841_1;
}
#if defined(__STDC_VERSION__)
int variant1841_2(int value, ...) {
    return value + 90;
}
#else
int variant1841_2(int value) {
    return value + 90;
}
#endif
struct Node1841_3 {
    int value;
    struct Node1841_3 *next;
};

int length1841_3(const struct Node1841_3 *head) {
    int count = 0;
    const struct Node1841_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
typedef int (*predicate1841_4)(int value);

typedef struct Slot1841_4 {
    int slots[16];
    predicate1841_4 accept;
} Slot1841_4;

int slot_accept1841_4(Slot1841_4 *store, int value) {
    return store->accept(value) + 7;
}
#if defined(__STDC_VERSION__)
int variant1841_5(int value, ...) {
    return value + 14;
}
#else
int variant1841_5(int value) {
    return value + 14;
}
#endif
struct Node1841_6 {
    int value;
    struct Node1841_6 *next;
};

int length1841_6(const struct Node1841_6 *head) {
    int count = 0;
    const struct Node1841_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
