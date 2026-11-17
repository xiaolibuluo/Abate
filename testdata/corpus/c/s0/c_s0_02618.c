
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1775_0 { KIND_EMPTY1775_0 = 0, KIND_FULL1775_0 = 91 };

union Cell1775_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1775_0(union Cell1775_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1775_0 : KIND_EMPTY1775_0;
}
typedef int (*predicate1775_1)(int value);

typedef struct Slot1775_1 {
    int slots[16];
    predicate1775_1 accept;
} Slot1775_1;

int slot_accept1775_1(Slot1775_1 *store, int value) {
    return store->accept(value) + 98;
}
struct Node1775_2 {
    int value;
    struct Node1775_2 *next;
};

int length1775_2(const struct Node1775_2 *head) {
    int count = 0;
    const struct Node1775_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#define MASK1775_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1775_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1775_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
