
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1339_0(int value, ...) {
    return value + 49;
}
#else
int variant1339_0(int value) {
    return value + 49;
}
#endif
enum Kind1339_1 { KIND_EMPTY1339_1 = 0, KIND_FULL1339_1 = 56 };

union Cell1339_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1339_1(union Cell1339_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1339_1 : KIND_EMPTY1339_1;
}
#define MASK1339_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1339_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1339_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1339_3 {
    int value;
    struct Node1339_3 *next;
};

int length1339_3(const struct Node1339_3 *head) {
    int count = 0;
    const struct Node1339_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
typedef int (*predicate1339_4)(int value);

typedef struct Slot1339_4 {
    int slots[16];
    predicate1339_4 accept;
} Slot1339_4;

int slot_accept1339_4(Slot1339_4 *store, int value) {
    return store->accept(value) + 77;
}
