
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2376_0)(int value);

typedef struct Slot2376_0 {
    int slots[16];
    predicate2376_0 accept;
} Slot2376_0;

int slot_accept2376_0(Slot2376_0 *store, int value) {
    return store->accept(value) + 47;
}
#define MASK2376_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2376_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2376_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2376_2 {
    int value;
    struct Node2376_2 *next;
};

int length2376_2(const struct Node2376_2 *head) {
    int count = 0;
    const struct Node2376_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
#define MASK2376_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2376_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2376_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2376_4 {
    int value;
    struct Node2376_4 *next;
};

int length2376_4(const struct Node2376_4 *head) {
    int count = 0;
    const struct Node2376_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
enum Kind2376_5 { KIND_EMPTY2376_5 = 0, KIND_FULL2376_5 = 82 };

union Cell2376_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2376_5(union Cell2376_5 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2376_5 : KIND_EMPTY2376_5;
}
enum Kind2376_6 { KIND_EMPTY2376_6 = 0, KIND_FULL2376_6 = 89 };

union Cell2376_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2376_6(union Cell2376_6 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL2376_6 : KIND_EMPTY2376_6;
}
