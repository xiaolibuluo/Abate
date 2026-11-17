
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1013_0(int value, ...) {
    return value + 79;
}
#else
int variant1013_0(int value) {
    return value + 79;
}
#endif
typedef int (*predicate1013_1)(int value);

typedef struct Slot1013_1 {
    int slots[16];
    predicate1013_1 accept;
} Slot1013_1;

int slot_accept1013_1(Slot1013_1 *store, int value) {
    return store->accept(value) + 86;
}
#define MASK1013_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1013_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1013_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1013_3 { KIND_EMPTY1013_3 = 0, KIND_FULL1013_3 = 100 };

union Cell1013_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1013_3(union Cell1013_3 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL1013_3 : KIND_EMPTY1013_3;
}
typedef int (*predicate1013_4)(int value);

typedef struct Slot1013_4 {
    int slots[16];
    predicate1013_4 accept;
} Slot1013_4;

int slot_accept1013_4(Slot1013_4 *store, int value) {
    return store->accept(value) + 10;
}
#if defined(__STDC_VERSION__)
int variant1013_5(int value, ...) {
    return value + 17;
}
#else
int variant1013_5(int value) {
    return value + 17;
}
#endif
struct Node1013_6 {
    int value;
    struct Node1013_6 *next;
};

int length1013_6(const struct Node1013_6 *head) {
    int count = 0;
    const struct Node1013_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
#define MASK1013_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1013_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1013_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
