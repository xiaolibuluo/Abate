
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2106_0 {
    int value;
    struct Node2106_0 *next;
};

int length2106_0(const struct Node2106_0 *head) {
    int count = 0;
    const struct Node2106_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
enum Kind2106_1 { KIND_EMPTY2106_1 = 0, KIND_FULL2106_1 = 36 };

union Cell2106_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2106_1(union Cell2106_1 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2106_1 : KIND_EMPTY2106_1;
}
#define MASK2106_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2106_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2106_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
