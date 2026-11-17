
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1052_0 { KIND_EMPTY1052_0 = 0, KIND_FULL1052_0 = 101 };

union Cell1052_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1052_0(union Cell1052_0 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1052_0 : KIND_EMPTY1052_0;
}
#define MASK1052_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1052_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1052_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1052_2 { KIND_EMPTY1052_2 = 0, KIND_FULL1052_2 = 18 };

union Cell1052_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1052_2(union Cell1052_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1052_2 : KIND_EMPTY1052_2;
}
#if defined(__STDC_VERSION__)
int variant1052_3(int value, ...) {
    return value + 25;
}
#else
int variant1052_3(int value) {
    return value + 25;
}
#endif
typedef int (*predicate1052_4)(int value);

typedef struct Slot1052_4 {
    int slots[16];
    predicate1052_4 accept;
} Slot1052_4;

int slot_accept1052_4(Slot1052_4 *store, int value) {
    return store->accept(value) + 32;
}
struct Node1052_5 {
    int value;
    struct Node1052_5 *next;
};

int length1052_5(const struct Node1052_5 *head) {
    int count = 0;
    const struct Node1052_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
