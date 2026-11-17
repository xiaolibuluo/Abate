
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1681_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1681_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1681_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1681_1 {
    int value;
    struct Node1681_1 *next;
};

int length1681_1(const struct Node1681_1 *head) {
    int count = 0;
    const struct Node1681_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
typedef int (*predicate1681_2)(int value);

typedef struct Slot1681_2 {
    int slots[16];
    predicate1681_2 accept;
} Slot1681_2;

int slot_accept1681_2(Slot1681_2 *store, int value) {
    return store->accept(value) + 47;
}
#if defined(__STDC_VERSION__)
int variant1681_3(int value, ...) {
    return value + 54;
}
#else
int variant1681_3(int value) {
    return value + 54;
}
#endif
#if defined(__STDC_VERSION__)
int variant1681_4(int value, ...) {
    return value + 61;
}
#else
int variant1681_4(int value) {
    return value + 61;
}
#endif
enum Kind1681_5 { KIND_EMPTY1681_5 = 0, KIND_FULL1681_5 = 68 };

union Cell1681_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1681_5(union Cell1681_5 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1681_5 : KIND_EMPTY1681_5;
}
