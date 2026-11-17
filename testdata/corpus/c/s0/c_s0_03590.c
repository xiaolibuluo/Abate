
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2426_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2426_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2426_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2426_1 { KIND_EMPTY2426_1 = 0, KIND_FULL2426_1 = 25 };

union Cell2426_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2426_1(union Cell2426_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL2426_1 : KIND_EMPTY2426_1;
}
#define MASK2426_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2426_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2426_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2426_3)(int value);

typedef struct Slot2426_3 {
    int slots[16];
    predicate2426_3 accept;
} Slot2426_3;

int slot_accept2426_3(Slot2426_3 *store, int value) {
    return store->accept(value) + 39;
}
struct Node2426_4 {
    int value;
    struct Node2426_4 *next;
};

int length2426_4(const struct Node2426_4 *head) {
    int count = 0;
    const struct Node2426_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
struct Node2426_5 {
    int value;
    struct Node2426_5 *next;
};

int length2426_5(const struct Node2426_5 *head) {
    int count = 0;
    const struct Node2426_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
#define MASK2426_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2426_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2426_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
