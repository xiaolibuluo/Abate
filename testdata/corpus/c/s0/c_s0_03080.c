
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2081_0 {
    int value;
    struct Node2081_0 *next;
};

int length2081_0(const struct Node2081_0 *head) {
    int count = 0;
    const struct Node2081_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
#define MASK2081_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2081_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2081_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2081_2 { KIND_EMPTY2081_2 = 0, KIND_FULL2081_2 = 9 };

union Cell2081_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2081_2(union Cell2081_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2081_2 : KIND_EMPTY2081_2;
}
struct Node2081_3 {
    int value;
    struct Node2081_3 *next;
};

int length2081_3(const struct Node2081_3 *head) {
    int count = 0;
    const struct Node2081_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
typedef int (*predicate2081_4)(int value);

typedef struct Slot2081_4 {
    int slots[16];
    predicate2081_4 accept;
} Slot2081_4;

int slot_accept2081_4(Slot2081_4 *store, int value) {
    return store->accept(value) + 23;
}
typedef int (*predicate2081_5)(int value);

typedef struct Slot2081_5 {
    int slots[16];
    predicate2081_5 accept;
} Slot2081_5;

int slot_accept2081_5(Slot2081_5 *store, int value) {
    return store->accept(value) + 30;
}
