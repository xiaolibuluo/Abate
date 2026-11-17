
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1338_0 { KIND_EMPTY1338_0 = 0, KIND_FULL1338_0 = 36 };

union Cell1338_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1338_0(union Cell1338_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1338_0 : KIND_EMPTY1338_0;
}
struct Node1338_1 {
    int value;
    struct Node1338_1 *next;
};

int length1338_1(const struct Node1338_1 *head) {
    int count = 0;
    const struct Node1338_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
typedef int (*predicate1338_2)(int value);

typedef struct Slot1338_2 {
    int slots[16];
    predicate1338_2 accept;
} Slot1338_2;

int slot_accept1338_2(Slot1338_2 *store, int value) {
    return store->accept(value) + 50;
}
#define MASK1338_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1338_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1338_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
