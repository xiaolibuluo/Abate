
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2071_0)(int value);

typedef struct Slot2071_0 {
    int slots[16];
    predicate2071_0 accept;
} Slot2071_0;

int slot_accept2071_0(Slot2071_0 *store, int value) {
    return store->accept(value) + 59;
}
int scan2071_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 66;
}
enum Kind2071_2 { KIND_EMPTY2071_2 = 0, KIND_FULL2071_2 = 73 };

union Cell2071_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2071_2(union Cell2071_2 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2071_2 : KIND_EMPTY2071_2;
}
#define MASK2071_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2071_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2071_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2071_4 {
    int value;
    struct Node2071_4 *next;
};

int length2071_4(const struct Node2071_4 *head) {
    int count = 0;
    const struct Node2071_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
