
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2110_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2110_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2110_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2110_1 { KIND_EMPTY2110_1 = 0, KIND_FULL2110_1 = 88 };

union Cell2110_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2110_1(union Cell2110_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2110_1 : KIND_EMPTY2110_1;
}
struct Node2110_2 {
    int value;
    struct Node2110_2 *next;
};

int length2110_2(const struct Node2110_2 *head) {
    int count = 0;
    const struct Node2110_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
