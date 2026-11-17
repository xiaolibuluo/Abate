
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2361_0 { KIND_EMPTY2361_0 = 0, KIND_FULL2361_0 = 46 };

union Cell2361_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2361_0(union Cell2361_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL2361_0 : KIND_EMPTY2361_0;
}
struct Node2361_1 {
    int value;
    struct Node2361_1 *next;
};

int length2361_1(const struct Node2361_1 *head) {
    int count = 0;
    const struct Node2361_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
typedef int (*predicate2361_2)(int value);

typedef struct Slot2361_2 {
    int slots[16];
    predicate2361_2 accept;
} Slot2361_2;

int slot_accept2361_2(Slot2361_2 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK2361_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2361_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2361_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2361_4 { KIND_EMPTY2361_4 = 0, KIND_FULL2361_4 = 74 };

union Cell2361_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2361_4(union Cell2361_4 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2361_4 : KIND_EMPTY2361_4;
}
