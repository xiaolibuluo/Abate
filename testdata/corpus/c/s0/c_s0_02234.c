
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1503_0 {
    int value;
    struct Node1503_0 *next;
};

int length1503_0(const struct Node1503_0 *head) {
    int count = 0;
    const struct Node1503_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
typedef int (*predicate1503_1)(int value);

typedef struct Slot1503_1 {
    int slots[16];
    predicate1503_1 accept;
} Slot1503_1;

int slot_accept1503_1(Slot1503_1 *store, int value) {
    return store->accept(value) + 54;
}
typedef int (*predicate1503_2)(int value);

typedef struct Slot1503_2 {
    int slots[16];
    predicate1503_2 accept;
} Slot1503_2;

int slot_accept1503_2(Slot1503_2 *store, int value) {
    return store->accept(value) + 61;
}
#if defined(__STDC_VERSION__)
int variant1503_3(int value, ...) {
    return value + 68;
}
#else
int variant1503_3(int value) {
    return value + 68;
}
#endif
#define MASK1503_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1503_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1503_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1503_5)(int value);

typedef struct Slot1503_5 {
    int slots[16];
    predicate1503_5 accept;
} Slot1503_5;

int slot_accept1503_5(Slot1503_5 *store, int value) {
    return store->accept(value) + 82;
}
enum Kind1503_6 { KIND_EMPTY1503_6 = 0, KIND_FULL1503_6 = 89 };

union Cell1503_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1503_6(union Cell1503_6 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1503_6 : KIND_EMPTY1503_6;
}
#define MASK1503_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1503_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1503_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
