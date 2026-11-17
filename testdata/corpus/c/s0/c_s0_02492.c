
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1694_0)(int value);

typedef struct Slot1694_0 {
    int slots[16];
    predicate1694_0 accept;
} Slot1694_0;

int slot_accept1694_0(Slot1694_0 *store, int value) {
    return store->accept(value) + 8;
}
enum Kind1694_1 { KIND_EMPTY1694_1 = 0, KIND_FULL1694_1 = 15 };

union Cell1694_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1694_1(union Cell1694_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1694_1 : KIND_EMPTY1694_1;
}
struct Node1694_2 {
    int value;
    struct Node1694_2 *next;
};

int length1694_2(const struct Node1694_2 *head) {
    int count = 0;
    const struct Node1694_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
#define MASK1694_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1694_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1694_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1694_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1694_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1694_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1694_5)(int value);

typedef struct Slot1694_5 {
    int slots[16];
    predicate1694_5 accept;
} Slot1694_5;

int slot_accept1694_5(Slot1694_5 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant1694_6(int value, ...) {
    return value + 50;
}
#else
int variant1694_6(int value) {
    return value + 50;
}
#endif
