
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1626_0 {
    int value;
    struct Node1626_0 *next;
};

int length1626_0(const struct Node1626_0 *head) {
    int count = 0;
    const struct Node1626_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
enum Kind1626_1 { KIND_EMPTY1626_1 = 0, KIND_FULL1626_1 = 101 };

union Cell1626_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1626_1(union Cell1626_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1626_1 : KIND_EMPTY1626_1;
}
typedef int (*predicate1626_2)(int value);

typedef struct Slot1626_2 {
    int slots[16];
    predicate1626_2 accept;
} Slot1626_2;

int slot_accept1626_2(Slot1626_2 *store, int value) {
    return store->accept(value) + 11;
}
#define MASK1626_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1626_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1626_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
