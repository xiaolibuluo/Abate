
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2116_0)(int value);

typedef struct Slot2116_0 {
    int slots[16];
    predicate2116_0 accept;
} Slot2116_0;

int slot_accept2116_0(Slot2116_0 *store, int value) {
    return store->accept(value) + 62;
}
#define MASK2116_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2116_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2116_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2116_2 { KIND_EMPTY2116_2 = 0, KIND_FULL2116_2 = 76 };

union Cell2116_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2116_2(union Cell2116_2 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL2116_2 : KIND_EMPTY2116_2;
}
struct Node2116_3 {
    int value;
    struct Node2116_3 *next;
};

int length2116_3(const struct Node2116_3 *head) {
    int count = 0;
    const struct Node2116_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
