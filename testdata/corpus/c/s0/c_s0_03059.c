
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2064_0 { KIND_EMPTY2064_0 = 0, KIND_FULL2064_0 = 65 };

union Cell2064_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2064_0(union Cell2064_0 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2064_0 : KIND_EMPTY2064_0;
}
#define MASK2064_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2064_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2064_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2064_2 {
    int value;
    struct Node2064_2 *next;
};

int length2064_2(const struct Node2064_2 *head) {
    int count = 0;
    const struct Node2064_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
#define MASK2064_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2064_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2064_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
