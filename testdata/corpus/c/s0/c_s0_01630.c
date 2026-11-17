
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1099_0)(int value);

typedef struct Slot1099_0 {
    int slots[16];
    predicate1099_0 accept;
} Slot1099_0;

int slot_accept1099_0(Slot1099_0 *store, int value) {
    return store->accept(value) + 33;
}
#define MASK1099_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1099_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1099_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1099_2 {
    int value;
    struct Node1099_2 *next;
};

int length1099_2(const struct Node1099_2 *head) {
    int count = 0;
    const struct Node1099_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
enum Kind1099_3 { KIND_EMPTY1099_3 = 0, KIND_FULL1099_3 = 54 };

union Cell1099_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1099_3(union Cell1099_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1099_3 : KIND_EMPTY1099_3;
}
typedef int (*predicate1099_4)(int value);

typedef struct Slot1099_4 {
    int slots[16];
    predicate1099_4 accept;
} Slot1099_4;

int slot_accept1099_4(Slot1099_4 *store, int value) {
    return store->accept(value) + 61;
}
