
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1641_0 { KIND_EMPTY1641_0 = 0, KIND_FULL1641_0 = 95 };

union Cell1641_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1641_0(union Cell1641_0 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1641_0 : KIND_EMPTY1641_0;
}
#define MASK1641_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1641_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1641_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1641_2 { KIND_EMPTY1641_2 = 0, KIND_FULL1641_2 = 12 };

union Cell1641_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1641_2(union Cell1641_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL1641_2 : KIND_EMPTY1641_2;
}
struct Node1641_3 {
    int value;
    struct Node1641_3 *next;
};

int length1641_3(const struct Node1641_3 *head) {
    int count = 0;
    const struct Node1641_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
#if defined(__STDC_VERSION__)
int variant1641_4(int value, ...) {
    return value + 26;
}
#else
int variant1641_4(int value) {
    return value + 26;
}
#endif
typedef int (*predicate1641_5)(int value);

typedef struct Slot1641_5 {
    int slots[16];
    predicate1641_5 accept;
} Slot1641_5;

int slot_accept1641_5(Slot1641_5 *store, int value) {
    return store->accept(value) + 33;
}
